package post

import (
	"context"
	"errors"
	"fmt"
	"socialNetwork/entity"
)

// func (p *post) CreatePostRepo(post entity.Post, id int) (int, int, error) {
// 	return 0, http.StatusCreated, nil
// }

// func (p *post) CreateGroup(post entity.Post, id int) (int, int, error) {
// 	return 0, http.StatusCreated, nil
// }

func (p *post) CanSeePost(userid, postid int) bool {
	prep, err := p.db.Prepare(`SELECT
		    1
		FROM posts AS post 
		INNER JOIN users AS user ON user.id = post.user_id
		LEFT JOIN group_members AS gm ON gm.group_id=post.group_id AND gm.member_id = $1
		LEFT JOIN follows AS follow ON follow.followed_id = post.user_id AND follow.follower_id = 1
		WHERE
		    (post.group_id IS NOT NULL AND gm.member_id = $1 AND post.id=$2)
		    OR 
		    (post.group_id IS NULL AND follow.follower_id = $1 AND post.id=$2);`)
	if err != nil {
		return false
	}
	row := prep.QueryRow(userid, postid)
	var res bool
	err = row.Scan(&res)
	if err != nil {
		return false
	}
	return res
}

func (p *post) GetPostsByUserID(ctx context.Context, userID int, limit int, offset int) ([]entity.Post, error) {
	query := `
		SELECT id, title, image, content, user_id, status, group_id,  created_at, updated_at
		FROM posts
		WHERE user_id = ?
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error preparing query in GetGroupsByUserID function: %v", err))
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Image,
			&post.Content,
			&post.UserID,
			&post.Status,
			&post.GroupID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
