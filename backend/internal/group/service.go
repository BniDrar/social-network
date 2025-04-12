package group

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"socialNetwork/entity"
)

func (g *group) GetGroupsByUserService(ctx context.Context, limit, offset int) (entity.Groups, error) {
	userID, ok := ctx.Value(entity.ContextID).(int)
	if !ok {
		return nil, errors.New("user id not found in context")
	}
	groups, err := g.GetGroupsByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (g *group) GetGroupByIdService(ctx context.Context, groupID int) (entity.Group, error) {
	userID, ok := ctx.Value(entity.ContextID).(int)
	if !ok {
		return entity.Group{}, errors.New("user id not found in context")
	}
	group, err := g.GetGroupByIdRepository(ctx, userID, groupID)
	if err != nil {
		return entity.Group{}, err
	}
	return group, nil
}

func (g *group) CreateGroupService(ctx context.Context, group entity.Group) (entity.Group, error) {
	group, err := g.CreateGroupRepository(ctx, group)
	if err != nil {
		return entity.Group{}, err
	}
	return group, nil
}

func (g *group) GetAllGroupsService(ctx context.Context, limit, offset, typeGroup int) (entity.Groups, error) {
	groups, err := g.GetAllGroupsRepository(ctx, limit, offset, typeGroup)
	if err != nil {
		return nil, err
	}
	return groups, nil
}


//----------------events---------------------------------
func (g *group) CreateEventService(ctx context.Context, event entity.Event) (int, int, error) {
	// check if the user is a member of the group
	fmt.Println(event.GroupID)
	isMember, err := g.IsMemberRepository(ctx, event.GroupID)
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
	return eventId, status, nil
}

func (g *group) GetEventService(ctx context.Context, eventID int) (entity.Event, int, error) {
	event, err := g.GetEventRepository(ctx, eventID)
	if err != nil {
		return entity.Event{}, http.StatusInternalServerError, err
	}
	isMember, err := g.IsMemberRepository(ctx, event.GroupID)
	if err != nil || !isMember {
		if err == nil {
			err = errors.New("user is not a member of the group")
		}
		return entity.Event{}, http.StatusBadRequest,err
	}
	return event, http.StatusOK,nil
}

func (g *group) VoteEventService(ctx context.Context, vote entity.Engagement) (int, int, error) {
	// check if the user is a member of the group
	event, err := g.GetEventRepository(ctx, vote.EventID)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	fmt.Println(event.UserID, event.GroupID)
	isMember, err := g.IsMemberRepository(ctx, event.GroupID)
	if err != nil || !isMember {
		if err == nil {
			err = errors.New("user is not a member of the group")
		}
		return 0, http.StatusBadRequest, err
	}
	eventId, status, err := g.VoteEventRepository(ctx, event.ID, vote.Status)
	if err != nil {
		return eventId, status, err
	}

	return eventId, status, nil
}
