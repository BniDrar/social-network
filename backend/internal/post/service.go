package post

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"socialNetwork/entity"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func (p *post) Service_GetAll(ctx context.Context) (data []byte, err error) {
	id := 1 // ctx.Value(entity.ContextID).(int)
	posts, err := p.Repo_GetAll(ctx, id)
	if err != nil {
		return
	}
	data, err = json.Marshal(posts)
	return
}

func (p *post) Service_GetOne(ctx context.Context, post_str string) (data []byte, err error) {
	id := 1 // ctx.Value(entity.ContextID).(int)
	post_id, err := strconv.Atoi(post_str)
	if err != nil {
		return
	}
	if p.Repo_UserCanPost(ctx, id, post_id) {
		var post entity.Post
		post, err = p.Repo_GetOne(ctx, post_id)
		if err != nil {
			return
		}
		data, err = json.Marshal(post)
		return
	} else {
		err = errors.New("you can't see post")
		return
	}
}

func (p *post) Service_CreateOne(ctx context.Context, body io.ReadCloser, users []string, post entity.Post) error {
	id := 1 // ctx.Value(entity.ContextID).(int)
	ImageFileName := post.Image
	json.NewDecoder(body).Decode(&post)
	post.Image = ImageFileName
	err := post.Validate(users)
	if err != nil {
		return errors.New(string("{ error: " + err.Error() + " }"))
	}
	p.Repo_CreatePost(ctx, id, post)
	return nil
}

func (p *post) Service_React(ctx context.Context, body io.ReadCloser) (err error) {
	id := 1 // ctx.Value(entity.ContextID).(int)
	react := entity.Vote{}
	json.NewDecoder(body).Decode(&react)
	if p.Repo_UserCanPost(ctx, id, react.ID) {
		err = p.Repo_React(ctx, id, react)
	} else {
		err = errors.New("you can't see the post")
	}
	return
}

func (p *post) GetPostsByUserService(ctx context.Context, username string) (data []byte, err error) {
	userID := 1 // ctx.Value(entity.ContextID).(int)
	posts, err := p.GetPostsByUserID(ctx, userID, username)
	if err != nil {
		return
	}
	data, err = json.Marshal(posts)
	if err != nil {
		return
	}
	return
}
