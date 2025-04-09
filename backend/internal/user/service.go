package user

import (
	"context"
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
	_, err = s.GetUserByUsername(user.Nickname)
	if err != nil {
		return http.StatusBadRequest, errors.New("nickname already exists")
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

func (u *user) LogoutService(user entity.User) error {
	// do something
	return nil
}

func (u *user) authenticateService(email, password string) (int, error) {
	return u.authenticateRepo(email, password)
}

func (u *user) UserProfile(ctx context.Context, nickname string) (int, entity.User, error) {
	user, err := u.GetUserByUsername(nickname)
	if err != nil {
		return http.StatusInternalServerError, user, fmt.Errorf("erro while getting the profile from the database, err: %v", err)
	}
	user.Password = ""
	if user.Status == entity.PublicUser {
		return http.StatusOK, user, nil
	}
	userId := ctx.Value(entity.ContextID).(int)
	exists, err := u.isFollowedBy(userId, int(user.ID))
	if err != nil && exists && userId == int(user.ID) {
		return http.StatusInternalServerError, user, err
	}
	return http.StatusOK, user, nil
}

func (u *user) FollowService(user entity.User) error {
	// do something
	return nil
}

func (u *user) FollowersService(user entity.User) error {
	// do something
	return nil
}

// func (s *user) FollowersService(user entity.User) error {
// 	// do something
// 	return nil
// }

// func (s *user) IsExistsService(user uint) bool {
// 	// do something
// 	return false
// }

// func (s *user) DeleteUserService(user entity.User) error {
// 	return nil
// }
