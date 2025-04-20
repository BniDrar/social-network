package comment

import (
	"context"
	"errors"
	"net/http"

	"socialNetwork/entity"
)

func (c *comment) GetCommentsService(ctx context.Context, post entity.Post) (int, []entity.Comment, error) {
	if !c.CanSeePost(ctx.Value(entity.ContextID).(int), post.ID) {
		return http.StatusForbidden, nil, errors.New("you don't allowed for this action")
	}
	return c.getPostComments(ctx, post.ID)
}

func (c *comment) CreateCommentService(ctx context.Context, commnt entity.Comment) (int, int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	if !c.CanSeePost(userId, commnt.PostID) {
		return 0, http.StatusForbidden, errors.New("you don't allowed for this action")
	}
	commnt.UserID = userId
	commentId, status, err := c.CreateCommentRepo(ctx, commnt)
	if err != nil {
		if status != http.StatusBadRequest {
			err = nil
		}
		return 0, status, err
	}
	return commentId, http.StatusCreated, nil
}
