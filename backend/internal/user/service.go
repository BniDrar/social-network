package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/utils"
)

func (s *user) RegisterService(user entity.User) (int, error) {
	// check credentials
	if err := utils.ValidateRegisterCredentials(user); err != nil {
		return http.StatusBadRequest, err
	}
	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	user.Password = hashedPassword
	// chek if user Email already exists
	_, err = s.GetUserByUsername(user.Email)
	if err != nil {
		return http.StatusBadRequest, errors.New("email already exists")
	}
	// check if user Nickname already exists
	if user.Nickname.String != "" {
		user.Nickname.Valid = true
		_, err = s.GetUserByUsername(user.Nickname.String)
		if err != nil {
			return http.StatusBadRequest, errors.New("nickname already exists")
		}
	}
	// save user to db
	err = s.CreateUser(user)
	if err != nil {
		if errors.Is(err, config.ErrUserAlreadyExists) {
			return http.StatusBadRequest, err
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (u *user) authenticateService(email, password string) (int, error) {
	return u.authenticateRepo(email, password)
}

func (u *user) UserProfile(ctx context.Context, targetId int) (int, entity.User, error) {
	user, err := u.GetUserProfileById(ctx, targetId)
	if err != nil {
		return http.StatusInternalServerError, user, fmt.Errorf("erro while getting the profile from the database, err: %v", err)
	}
	user.ProfileOwner = int(user.ID) == ctx.Value(entity.ContextID).(int)
	if user.Status == entity.PublicUser || user.ProfileOwner {
		return http.StatusOK, user, nil
	}
	userId := ctx.Value(entity.ContextID).(int)
	exists, err := u.isFollowedBy(userId, int(user.ID))
	if err != nil || !exists {
		if err == nil {
			return http.StatusUnauthorized, entity.User{}, errors.New("you can't access to the user profile")
		}
		return http.StatusInternalServerError, entity.User{}, err
	}
	return http.StatusOK, user, nil
}

func (u *user) GetUserPostsService(ctx context.Context, cursor entity.Cursor) ([]entity.Post, int, error) {
	id := ctx.Value(entity.ContextID).(int)
	user, err := u.GetUserProfileById(ctx, cursor.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if user.Status == entity.PrivateUser && id != cursor.ID {
		authorized, err := u.isFollowedBy(id, cursor.ID)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		if !authorized {
			return nil, http.StatusUnauthorized, errors.New("you can't access to the user profile")
		}
	}
	fmt.Println(cursor)
	posts, err := u.GetUserPostsRep(ctx, cursor)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return posts, http.StatusOK, nil
}

func (u *user) ChangeStatusService(ctx context.Context) error {
	id := ctx.Value(entity.ContextID).(int)
	return u.changeStatusRepo(ctx, id)
}

func (u *user) FollowersAndFollowedService(ctx context.Context, id int) (int, entity.Follows, error) {
	var (
		follows entity.Follows
		err     error
	)
	follows.Followers, err = u.GetFollowers(ctx, id)
	if err != nil {
		u.loger.Error.Println(err)
		return http.StatusInternalServerError, follows, errors.New("invalid user id")
	}
	utils.SetOnlineStatus(follows.Followers, u.hub.IsOnline)
	follows.Following, err = u.GetFollowing(ctx, id)
	if err != nil {
		u.loger.Error.Println(err)
		return http.StatusInternalServerError, follows, err
	}
	utils.SetOnlineStatus(follows.Following, u.hub.IsOnline)

	return http.StatusOK, follows, nil
}

func (u *user) FollowService(ctx context.Context, followedID int) (int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	if userId == followedID {
		return http.StatusBadRequest, errors.New("you can't follow yourself")
	}
	//check if the followed account is private
	user, err := u.GetUserProfileById(ctx, followedID)
	if err != nil {
		return http.StatusBadRequest, errors.New("unavailable user")
	}
	if user.Status == entity.PrivateUser {
		// create notification in database
		_, err := u.CreateFollowNotification(ctx, userId, followedID)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		// Get follower information
		follower, err := u.GetUserProfileById(ctx, userId)
		if err != nil {
			u.loger.Error.Printf("Error getting follower info: %v", err)
			return http.StatusInternalServerError, err
		}

		// Create notification message
		notification := entity.Notification{
			Type:       entity.FollowingNotification,
			SenderId:   userId,
			ReceiverID: followedID,
			Message:    fmt.Sprintf("%s requested to follow you", follower.Nickname),
		}
		notificationBytes, err := json.Marshal(notification)
		if err != nil {
			u.loger.Error.Printf("Error marshaling notification: %v", err)
			return http.StatusInternalServerError, err
		}

		// Send notification to the followed user
		u.hub.SendNotification(uint(followedID), notificationBytes)
		u.loger.Info.Printf("Sent follow request notification to user %d", followedID)

		return http.StatusOK, nil
	}
	notify, err := u.FollowRepository(userId, followedID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if notify {
		// Get follower information
		follower, err := u.GetUserProfileById(ctx, userId)
		if err != nil {
			u.loger.Error.Printf("Error getting follower info: %v", err)
			return http.StatusOK, nil // Return success even if notification fails
		}

		// Create notification message
		notification := entity.Notification{
			Type:       entity.FollowingNotification,
			SenderId:   userId,
			ReceiverID: followedID,
			Message:    fmt.Sprintf("%s started following you", follower.Nickname),
		}
		notificationBytes, err := json.Marshal(notification)
		if err != nil {
			u.loger.Error.Printf("Error marshaling notification: %v", err)
			return http.StatusOK, nil
		}

		// Send notification to the followed user
		u.hub.SendNotification(uint(followedID), notificationBytes)
		u.loger.Info.Printf("Sent follow notification to user %d", followedID)
	}

	return http.StatusOK, nil
}

func (u *user) FollowersService(ctx context.Context, followedId int) error {
	// do something
	return nil
}

func (u *user) processRequestResponse(ctx context.Context, notification entity.Notification) (int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	// this user is contained in the group
	if notification.ReceiverID != userId {

		return http.StatusForbidden, errors.New("forbidden access to this action")
	}
	if notification.Accepted {
		_, err := u.FollowRepository(notification.SenderId, userId)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}

func (u *user) userNotificationSerice(ctx context.Context) ([]entity.Notification, int, error) {
	return u.userNotificationRepo(ctx)
}
