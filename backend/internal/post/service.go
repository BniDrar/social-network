package post

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"socialNetwork/entity"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func (p *post) GetPostsService(ctx context.Context, limit, offset int) ([]entity.Post, int, error) {
	id := ctx.Value(entity.ContextID).(int)
	return p.Repo_GetAll(ctx, id, limit, offset)
}

func (p *post) GetPostService(ctx context.Context, postID int) (entity.Post, int, error) {
	id := ctx.Value(entity.ContextID).(int)
	if !p.Repo_UserCanPost(ctx, id, postID) {
		return entity.Post{}, http.StatusUnauthorized, errors.New("you don't have the permision")
	}
	return p.GetPostRepo(ctx, postID)
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

func (p *post) PostEngagementService(ctx context.Context, engagement entity.Engagement) (int, error) {
	engagement.UserID = ctx.Value(entity.ContextID).(int)
	if !p.Repo_UserCanPost(ctx, engagement.UserID, engagement.PostID) {
		return http.StatusUnauthorized, errors.New("you don't have the permission to react")
	}
	return p.PostEngagementRepo(ctx, engagement)
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
