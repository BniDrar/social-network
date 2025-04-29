package post

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"

	"socialNetwork/entity"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func (p *post) Service_GetAll(ctx context.Context) (data []byte, err error) {
	id := ctx.Value(entity.ContextID).(int)
	posts, err := p.Repo_GetAll(ctx, id)
	if err != nil {
		return
	}
	data, err = json.Marshal(posts)
	return
}

func (p *post) Service_GetOne(ctx context.Context, post_str string) (data []byte, err error) {
	id := ctx.Value(entity.ContextID).(int)
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

func (p *post) CreatePostService(ctx context.Context, post entity.Post, imageContent []byte) (int, int, error) {
	postID, status, err := p.SavePost(ctx, ctx.Value(entity.ContextID).(int), post)
	if err != nil {
		return 0, status, err
	}
	err = os.WriteFile(post.Image, imageContent, 0o444)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	return postID, http.StatusCreated, nil
}

func (p *post) Service_React(ctx context.Context, body io.ReadCloser) (err error) {
	id := ctx.Value(entity.ContextID).(int)
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
	userID := ctx.Value(entity.ContextID).(int)
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
