package group

import (
	"context"
	"errors"
	"socialNetwork/entity"
)


func (g *group) GetGroupsByUser(ctx context.Context,limit, offset int) (entity.Groups, error) {
	userID, ok := ctx.Value(entity.ContextID).(int)
	if !ok {
		return nil, errors.New("user id not found in context")
	}
	groups, err := g.GetGroupsByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return groups, nil
}