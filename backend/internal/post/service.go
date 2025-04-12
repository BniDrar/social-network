package post

import (
	"context"
	"errors"
	"socialNetwork/entity"
)

// func (p *post) CreatePostService(post entity.Post, id int) (int, int, error) {
// 	var (
// 		status, GroupID int
// 		err error
// 	)
// 	switch post.Status {
// 	case 0:
// 		//repo => create group
// 		GroupID, status, err = p.CreateGroup(post, id)
// 		post.GroupID = GroupID
// 		if err != nil {
// 			return 0, status, err
// 		}
// 	default:
// 		if !(post.Status == 1 || post.Status == 2) {
// 			return 0, http.StatusBadRequest, errors.New("invalid post status")
// 		}
// 	}

// 	// repo => create post
// 	return p.CreatePostRepo(post, id)
// }

// get
// 0 => table(group) contains user id
// 1 => table(follows) contains user id
// 2 => get direct

// react
// 0 => table(group) contains user id
// 1 => table(follows) contains user id
// if not status forbidden

func (p *post) GetPostsByUserService(ctx context.Context, limit, offset int) (entity.Groups, error) {
	userID, ok := ctx.Value(entity.ContextID).(int)
	if !ok {
		return nil, errors.New("user id not found in context")
	}
	groups, err := p.GetPostsByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
