package comment

import (
	"context"
	"errors"
	"net/http"
	"os"

	"socialNetwork/entity"
)

func (c *comment) GetCommentsService(ctx context.Context, postId int) (int, []entity.Comment, error) {
	if !c.CanSeePost(ctx.Value(entity.ContextID).(int), postId) {
		return http.StatusForbidden, nil, errors.New("you don't allowed for this action")
	}

	return c.getPostComments(ctx, postId)
}

func (c *comment) CreateCommentService(ctx context.Context, commnt entity.Comment, imageContent []byte) (int, int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	if !c.CanSeePost(userId, commnt.PostID) {
		return 0, http.StatusForbidden, errors.New("you don't allowed for this action")
	}
	commnt.UserID = userId
	commentId, status, err := c.CreateCommentRepo(ctx, commnt)
	if err != nil {
		return 0, status, err
	}
	if commnt.Image.Valid {
		err = os.WriteFile(commnt.Image.NullString.String, imageContent, 0o444)
		if err != nil {
			return 0, http.StatusInternalServerError, err
		}
	}
	return commentId, http.StatusCreated, nil
}
