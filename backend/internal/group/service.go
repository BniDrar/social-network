package group

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/utils"
)

func (g *group) GetGroupsByUserService(ctx context.Context, limit, offset int) (entity.Groups, error) {
	return g.GetGroupsByUserID(ctx, ctx.Value(entity.ContextID).(int), limit, offset)
}

// GetGroupByIdService returns a group by its ID.
func (g *group) GetGroupByIdService(ctx context.Context, groupID int) (int, entity.Group, error) {
	userID := ctx.Value(entity.ContextID).(int)
	exist, err := g.IsMemberRepository(ctx, userID, groupID)
	if err != nil || !exist {
		return http.StatusForbidden, entity.Group{}, errors.New("you can't access this group")
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

func (g *group) GetUsersThatCanJoinGroupService(ctx context.Context, groupID int) ([]entity.User, int, error) {
	exist, err := g.IsMemberRepository(ctx, ctx.Value(entity.ContextID).(int), groupID)
	if err != nil || !exist {
		return nil, http.StatusForbidden, errors.New("you are not allowed to that")
	}
	users, err := g.GetUsersThatCanJoinGroupRepository(ctx, groupID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return users, http.StatusOK, nil
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

	// Get group information
	group, err := g.GetGroupByIdRepository(ctx, invitation.InviterID, invitation.GroupId)
	if err != nil {
		g.loger.Error.Printf("Error getting group info: %v", err)
		return http.StatusInternalServerError, err
	}

	// Get inviter information from group members
	groupMembers, err := g.GetGroupMembersRepository(ctx, invitation.GroupId)
	if err != nil {
		g.loger.Error.Printf("Error getting group members: %v", err)
		return http.StatusInternalServerError, err
	}

	// Find the inviter in group members
	var inviter entity.User
	for _, member := range groupMembers {
		if member.ID == uint(invitation.InviterID) {
			inviter = member
			break
		}
	}
	if inviter.ID == 0 {
		inviter, err = g.GetUserByIdRepository(ctx, invitation.InviterID)
		if err != nil {
			g.loger.Error.Printf("Error getting inviter info: %v", err)
			return http.StatusInternalServerError, err
		}
	}
	notificationID, status, err := g.CreateInvitationNotification(ctx, invitation)
	if err != nil {
		return status, err
	}
	// Create notification message
	notification := entity.Notification{
		Id:         notificationID,
		Type:       entity.GroupInvitationNotification,
		GroupId:    invitation.GroupId,
		SenderId:   invitation.InviterID,
		ReceiverID: invitation.InvitedID,
		Message:    fmt.Sprintf("%s invited you to join the group '%s'", inviter.Nickname.String, group.Name),
	}
	notificationBytes, err := json.Marshal(notification)
	if err != nil {
		g.loger.Error.Printf("Error marshaling notification: %v", err)
		return http.StatusInternalServerError, err
	}

	// Send notification to the invited user
	g.Hub.SendNotification(uint(invitation.InvitedID), notificationBytes)
	g.loger.Info.Printf("Sent invitation notification to user %d", invitation.InvitedID)

	return http.StatusOK, nil
}

func (g *group) proccessInvitationResponse(ctx context.Context, notf entity.Notification) (int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	notification, err := g.getNotificationById(ctx, notf.Id)
	if err != nil {
		fmt.Println(err)
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
		status, err := g.addGroupMember(ctx, notification.SenderId, notf.GroupId)
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

	requester, err := g.GetUserByIdRepository(ctx, userId)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	// Create notification message
	notification := entity.Notification{
		Type:       entity.GroupParticipationNotification,
		GroupId:    group.ID,
		SenderId:   userId,
		ReceiverID: group.Admin,
		Message:    fmt.Sprintf("%s requested to join the group '%s'", requester.First, group.Name),
	}

	notificationID, status, err := g.CreateRequestJoiningNotification(ctx, notification)
	if err != nil {
		return 0, status, err
	}
	notification.Id = notificationID
	notificationBytes, err := json.Marshal(notification)
	if err != nil {
		g.loger.Error.Printf("Error marshaling notification: %v", err)
		return notificationID, http.StatusOK, nil // Return success even if notification fails
	}

	// Send notification to the group admin
	g.Hub.SendNotification(uint(group.Admin), notificationBytes)
	g.loger.Info.Printf("Sent join request notification to admin %d", group.Admin)

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
	userId := ctx.Value(entity.ContextID).(int)
	isMember, err := g.IsMemberRepository(ctx, userId, event.GroupID)
	if err != nil || !isMember {
		if err == nil {
			err = errors.New("user is not a member of the group")
		}
		return 0, http.StatusBadRequest, err
	}
	if !utils.ValidateEventCredentials(event) {
		return 0, http.StatusBadRequest, errors.New("invalid credentials")
	}
	eventId, status, err := g.CreateEventRepository(ctx, event)
	if err != nil {
		return 0, status, err
	}
	event.ID = eventId
	_, status, err = g.CreateEventNotification(ctx, event)
	if err != nil {
		return 0, status, err
	}

	// Get all group members to send notification
	groupMembers, err := g.GetGroupMembersRepository(ctx, event.GroupID)
	if err != nil {
		g.loger.Error.Printf("Error getting group members: %v", err)
		return eventId, status, nil // Return success even if notification fails
	}

	// Create notification message
	notification := entity.Notification{
		Type:     entity.EventNotification,
		GroupId:  event.GroupID,
		EventID:  eventId,
		Message:  fmt.Sprintf("New event created: %s", event.Title),
		SenderId: userId,
	}
	notificationBytes, err := json.Marshal(notification)
	if err != nil {
		g.loger.Error.Printf("Error marshaling notification: %v", err)
		return eventId, status, nil
	}

	// Send notification to all group members
	for _, member := range groupMembers {
		if member.ID == uint(userId) {
			continue // Skip sending notification to the sender
		}
		g.Hub.SendNotification(uint(member.ID), notificationBytes)
	}
	g.loger.Info.Printf("Sent event notification to %d group members", len(groupMembers))

	return eventId, status, nil
}

func (g *group) GetEventService(ctx context.Context, eventID int) (entity.Event, int, error) {
	userID := ctx.Value(entity.ContextID).(int)
	event, err := g.GetEventRepository(ctx, eventID, userID)
	if err != nil {
		return entity.Event{}, http.StatusInternalServerError, err
	}
	isMember, err := g.IsMemberRepository(ctx, userID, event.GroupID)
	if err != nil || !isMember {
		if err == nil {
			err = errors.New("user is not a member of the group")
		}
		return entity.Event{}, http.StatusBadRequest, err
	}
	return event, http.StatusOK, nil
}

func (g *group) VoteEventService(ctx context.Context, vote entity.Engagement) (int, int, error) {
	userID := ctx.Value(entity.ContextID).(int)
	event, err := g.GetEventRepository(ctx, vote.EventID, userID)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	// check if the user is a member of the group
	isMember, err := g.IsMemberRepository(ctx, userID, event.GroupID)
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

func (g *group) GetGroupEventsService(ctx context.Context, groupID int) ([]entity.Event, int, error) {
	exist, err := g.IsMemberRepository(ctx, ctx.Value(entity.ContextID).(int), groupID)
	if err != nil || !exist {
		return nil, http.StatusForbidden, errors.New("you are not allowed to that")
	}
	events, err := g.GetEventsRepository(ctx, groupID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return events, http.StatusOK, nil
}
