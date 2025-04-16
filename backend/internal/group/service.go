package group

import (
	"context"
	"errors"
	"socialNetwork/entity"
)

// GetGroupsByUserService returns a list of groups that the user is a member of.
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

// GetGroupByIdService returns a group by its ID.
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
// CreateGroupService creates a new group.
func (g *group) CreateGroupService(ctx context.Context, group entity.Group) (entity.Group, error) {
	group, err := g.CreateGroupRepository(ctx, group)
	if err != nil {
		return entity.Group{}, err
	}
	return group, nil
}

// GetAllGroupsService returns a list of all groups.
func (g *group) GetAllGroupsService(ctx context.Context, limit, offset,typeGroup int) (entity.Groups, error) {
	groups, err := g.GetAllGroupsRepository(ctx, limit, offset ,typeGroup)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (g *group) GetGroupMembersService(ctx context.Context, groupID int) ([]entity.User, error) {
	group, err := g.GetGroupMembersRepository(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *group) RequestToJoinGroupService(ctx context.Context, groupID, userID int) error {
	err := g.RequestToJoinGroupRepository(ctx, groupID, userID)
	if err != nil {
		return err
	}
	return nil
}
func (g *group) AcceptRequestToJoinGroupService(ctx context.Context, groupID, userID int) error {
	err := g.AcceptRequestToJoinGroupRepository(ctx, groupID, userID)
	if err != nil {
		return err
	}
	return nil
}