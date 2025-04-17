package comment

import (
	"context"
	"encoding/json"
	"io"

	"socialNetwork/entity"
)

func (c *comment) ServiceGetComments(ctx context.Context, body io.ReadCloser) (err error) {
	UserID := r.Context().Value(entity.ContextID).(int)
	var Comment entity.Comment
	err = json.NewDecoder(body).Decode(&Comment)
	if err != nil {
		return
	}
	Comment.UserID = UserID
	// add get comments repo
}
