package group

import (
	"context"
	"errors"
	"socialNetwork/entity"
)


func (g *group) GetGroupsByUserService(ctx context.Context,limit, offset int) (entity.Groups, error) {
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
	userID, ok := ctx.Value(entity.ContextID).(int)
	if !ok {
		return entity.Group{}, errors.New("user id not found in context")
	}
	group.UserID = userID
	group, err := g.CreateGroupRepository(ctx, group)
	if err != nil {
		return entity.Group{}, err
	}
	return group, nil
}

func (g *group) UpdateGroupService(ctx context.Context, group entity.Group) (entity.Group, error) {
	userID, ok := ctx.Value(entity.ContextID).(int)
	if !ok {
		return entity.Group{}, errors.New("user id not found in context")
	}
	group.UserID = userID
	group, err := g.UpdateGroupRepsitory(ctx, group)
	if err != nil {
		return entity.Group{}, err
	}
	return group, nil
}