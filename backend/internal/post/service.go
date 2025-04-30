package post

import (
	"context"
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
	if post.Image.Valid {
		err = os.WriteFile(post.Image.NullString.String, imageContent, 0o444)
		if err != nil {
			return 0, http.StatusInternalServerError, err
		}
	}
	return postID, http.StatusCreated, nil
}

func (p *post) PostEngagementService(ctx context.Context, postId int) (int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	if !p.Repo_UserCanPost(ctx, userId, postId) {
		return http.StatusUnauthorized, errors.New("you don't have the permission to react")
	}
	err := p.PostEngagementRepo(ctx, userId, postId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
