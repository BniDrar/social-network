package post

import (
	"context"
	"net/http"
	"os"

	"socialNetwork/entity"
	"socialNetwork/pkg/errors"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func (p *post) GetPostsService(ctx context.Context, cursor entity.Cursor) ([]entity.Post, int, error) {
	id := ctx.Value(entity.ContextID).(int)
	posts, status, err := p.getAllPostsRepo(ctx, id, cursor)
	if err != nil {
		return nil, status, err
	}
	for _, post := range posts {
		if post.Image.String != "" {
			post.Image.SetValid(true)
		}
	}
	return posts, status, nil
}

func (p *post) getPostsByGroup(ctx context.Context, cursor entity.Cursor) (int, []entity.Post, error) {
	userId := ctx.Value(entity.ContextID).(int)
	if (cursor.LastId != nil && cursor.Time == nil) || (cursor.LastId == nil && cursor.Time != nil) {
		return http.StatusBadRequest, nil, errors.BadRequest("Both last_id and creation_time must be provided together", nil)
	}
	posts, err := p.getPostsByGroupId(ctx, userId, cursor)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.InternalServerError("Failed to get group posts", err)
	}
	return http.StatusOK, posts, nil
}

func (p *post) GetPostService(ctx context.Context, postID int) (entity.Post, int, error) {
	id := ctx.Value(entity.ContextID).(int)
	if !p.Repo_UserCanPost(ctx, id, postID) {
		return entity.Post{}, http.StatusUnauthorized, errors.Unauthorized("You don't have permission to view this post", nil)
	}

	return p.GetPostRepo(ctx, postID)
}

func (p *post) CreatePostService(ctx context.Context, post entity.Post, imageContent []byte) (int, int, error) {
	postID, status, err := p.SavePost(ctx, ctx.Value(entity.ContextID).(int), post)
	if err != nil {
		return 0, status, err
	}

	if post.Image.Valid {
		if err := os.WriteFile(post.Image.NullString.String, imageContent, 0o644); err != nil {
			return 0, http.StatusInternalServerError, errors.InternalServerError("Failed to save post image", err)
		}
	}
	return postID, http.StatusCreated, nil
}

func (p *post) PostEngagementService(ctx context.Context, postId int) (int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	if !p.Repo_UserCanPost(ctx, userId, postId) {
		return http.StatusUnauthorized, errors.Unauthorized("You don't have permission to react to this post", nil)
	}

	if err := p.PostEngagementRepo(ctx, userId, postId); err != nil {
		return http.StatusInternalServerError, errors.InternalServerError("Failed to process post engagement", err)
	}

	return http.StatusOK, nil
}
