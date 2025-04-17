package comment

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"socialNetwork/entity"
)

func (c *comment) ServiceGetComments(ctx context.Context, body io.ReadCloser) (err error) {
	UserID := ctx.Value(entity.ContextID).(int)
	var Comment entity.Comment
	err = json.NewDecoder(body).Decode(&Comment)
	if err != nil {
		return
	}
	Comment.UserID = UserID
	// add get comments repo
	return
}


func (c *comment) CreateCommentService(ctx context.Context, commnt entity.Comment) (int, int, error) {
	userId:= ctx.Value(entity.ContextID).(int)
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
