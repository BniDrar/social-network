package group

import (
	"context"
	"errors"
	"net/http"

	"socialNetwork/entity"
)

func (g *group) GetGroupsByUserService(ctx context.Context, limit, offset int) (entity.Groups, error) {
	return g.GetGroupsByUserID(ctx, ctx.Value(entity.ContextID).(int), limit, offset)
}

// GetGroupByIdService returns a group by its ID.
func (g *group) GetGroupByIdService(ctx context.Context, groupID int) (int, entity.Group, error) {
	userID := ctx.Value(entity.ContextID).(int)
	exist, err := g.IsMemberRepository(ctx, userID, groupID)
	if err != nil || !exist {
		return http.StatusForbidden, entity.Group{}, errors.New("You can't access this group")
	}
	group, err := g.GetGroupByIdRepository(ctx, userID, groupID)
	if err != nil {
		return http.StatusBadRequest, entity.Group{}, err
	}
	return http.StatusOK, group, nil
}

// CreateGroupService creates a new group.
func (g *group) CreateGroupService(ctx context.Context, group entity.Group) (entity.Group, error) {
	group.Admin = ctx.Value(entity.ContextID).(int)
	return g.CreateGroupRepository(ctx, group)
}

func (g *group) GetAllGroupsService(ctx context.Context, typeGroup int) (entity.Groups, error) {
	return g.GetAllGroupsRepository(ctx, typeGroup)
}

func (g *group) GetGroupMembersService(ctx context.Context, groupID int) ([]entity.User, error) {
	return g.GetGroupMembersRepository(ctx, groupID)
}

/*---------------notification related functions ---------------------*/
func (g *group) inviteToJoinGroupService(ctx context.Context, invitation entity.Invitation) (int, error) {
	invitation.InviterID = ctx.Value(entity.ContextID).(int)
	// check if the inviter is a group member
	inviterIsAGroupMember, err := g.IsMemberRepository(ctx, invitation.InviterID, invitation.GroupId)
	if err != nil || !inviterIsAGroupMember {
		if err == nil {
			err = errors.New("you are not a member in that group")
		}
		return http.StatusBadRequest, err
	}

	// check the invited is not a group member
	invitedIsAGroupMember, err := g.IsMemberRepository(ctx, invitation.InvitedID, invitation.GroupId)
	if err != nil || invitedIsAGroupMember {
		if err == nil {
			err = errors.New("the user already a member")
		}
		return http.StatusBadRequest, err
	}

	notfID, status, err := g.CreateInvitationNotification(ctx, invitation)
	if err != nil {
		return status, err
	}

	g.loger.Info.Println(notfID)
	// send the notification via websocket
	g.ws.SendMessage(uint(invitation.InvitedID), []byte("something"))
	return http.StatusOK, nil
}

func (g *group) proccessInvitationResponse(ctx context.Context, notf entity.Notification) (int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	notification, err := g.getNotificationById(ctx, notf.Id)
	if err != nil {
		return http.StatusBadRequest, errors.New("invalid credentials")
	}
	if userId != notification.ReceiverID {
		return http.StatusForbidden, errors.New("you doesn't have the right permitions")
	}
	err = g.RemoveNotificationById(ctx, notf.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if notf.Accepted {
		status, err := g.addGroupMember(ctx, notf.Id, notf.GroupId)
		if err != nil {
			return status, err
		}
	}
	return http.StatusOK, nil
}

func (g *group) requestToJoingGroupService(ctx context.Context, invitation entity.Invitation) (int, int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	invitedIsAGroupMember, err := g.IsMemberRepository(ctx, userId, invitation.GroupId)
	if err != nil || invitedIsAGroupMember {
		if err == nil {
			err = errors.New("you already a member of the group")
		}
		return 0, http.StatusBadRequest, err
	}

	group, err := g.GetGroupByIdRepository(ctx, userId, invitation.GroupId)
	if err != nil {
		return 0, http.StatusBadRequest, errors.New("invalid group id")
	}
	invitation.InvitedID = group.Admin
	notificationID, status, err := g.CreateRequestJoiningNotification(ctx, invitation)
	if err != nil {
		return 0, status, err
	}
	// now after getting group information send to the admin via websocket
	return notificationID, http.StatusOK, nil
}

func (g *group) processRequestToJoinResponse(ctx context.Context, notf entity.Notification) (int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	notification, err := g.getNotificationById(ctx, notf.Id)
	if err != nil || notification.ReceiverID != userId {
		return http.StatusBadRequest, errors.New("bad request")
	}

	group, err := g.GetGroupByIdRepository(ctx, userId, notf.GroupId)
	if err != nil || group.Admin != userId {
		return http.StatusForbidden, errors.New("you are not allowed to that")
	}
	err = g.RemoveNotificationById(ctx, notf.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if notf.Accepted {
		status, err := g.addGroupMember(ctx, notification.SenderId, notification.GroupId)
		if err != nil {
			return status, err
		}
	}
	return http.StatusOK, nil
}

// ----------------events---------------------------------
func (g *group) CreateEventService(ctx context.Context, event entity.Event) (int, int, error) {
	// check if the user is a member of the group
	isMember, err := g.IsMemberRepository(ctx, ctx.Value(entity.ContextID).(int), event.GroupID)
	if err != nil || !isMember {
		if err == nil {
			err = errors.New("user is not a member of the group")
		}
		return 0, http.StatusBadRequest, err
	}
	eventId, status, err := g.CreateEventRepository(ctx, event)
	if err != nil {
		return 0, status, err
	}
	notificationID, status, err := g.CreateEventNotification(ctx, event)
	if err != nil {
		return 0, status, err
	}
	// upstreat the notificationId and there information in the websocket to all the group members
	g.loger.Info.Println("the notification id is", notificationID)
	return eventId, status, nil
}

func (g *group) GetEventService(ctx context.Context, eventID int) (entity.Event, int, error) {
	event, err := g.GetEventRepository(ctx, eventID)
	if err != nil {
		return entity.Event{}, http.StatusInternalServerError, err
	}
	isMember, err := g.IsMemberRepository(ctx, ctx.Value(entity.ContextID).(int), event.GroupID)
	if err != nil || !isMember {
		if err == nil {
			err = errors.New("user is not a member of the group")
		}
		return entity.Event{}, http.StatusBadRequest, err
	}
	return event, http.StatusOK, nil
}

func (g *group) VoteEventService(ctx context.Context, vote entity.Engagement) (int, int, error) {
	// check if the user is a member of the group
	event, err := g.GetEventRepository(ctx, vote.EventID)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	isMember, err := g.IsMemberRepository(ctx, ctx.Value(entity.ContextID).(int), event.GroupID)
	if err != nil || !isMember {
		if err == nil {
			err = errors.New("user is not a member of the group")
		}
		return 0, http.StatusForbidden, err
	}
	eventId, status, err := g.VoteEventRepository(ctx, event.ID, vote.Status)
	if err != nil {
		return eventId, status, err
	}

	return eventId, status, nil
}
