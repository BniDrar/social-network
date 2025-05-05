package comment

import (
	"context"
	"errors"
	"net/http"

	"socialNetwork/entity"
)

func (c *comment) CanSeePost(userid, postid int) bool {
	prep, err := c.db.Prepare(`SELECT
		    post.id
		FROM
			posts AS post
		LEFT JOIN follows AS follow ON follow.followed_id = post.user_id
			AND follow.follower_id = $1
		LEFT JOIN group_members AS gm ON gm.group_id = post.group_id
			AND gm.member_id = $1
		WHERE
		    (post.status = 0 AND gm.member_id IS NOT NULL AND post.id = $2)
		    OR (post.status = 1 AND follow.follower_id = $1 AND post.id = $2)
			OR (post.status = 2 AND post.id = $2)
			OR (post.user_id = $1);
		ORDER BY post.created_at DESC`)
	if err != nil {
		return false
	}
	row := prep.QueryRow(userid, postid)
	var res int
	err = row.Scan(&res)
	if err != nil {
		return false
	}
	return true
}

func (c *comment) CreateCommentRepo(ctx context.Context, commnt entity.Comment) (int, int, error) {
	query := `INSERT INTO comments (post_id, user_id, content, image)
	VALUES (?, ?, ?, ?)`
	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	res, err := stmt.ExecContext(ctx, commnt.PostID, commnt.UserID, commnt.Content, commnt.Image)
	if err != nil {
		return 0, http.StatusBadRequest, errors.New("invalid credentials")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	return int(id), http.StatusCreated, nil
}

func (c *comment) getPostComments(ctx context.Context, postId int) (int, []entity.Comment, error) {
	query := `
		SELECT 
			comments.id,
			users.nickname AS creater_name,
			users.avatar AS avatar,
			comments.post_id,
			comments.user_id,
			comments.content,
			comments.image
		FROM comments
		JOIN users ON users.id = comments.user_id
		WHERE comments.post_id = $1
		ORDER BY comments.id ASC
	`

	rows, err := c.db.QueryContext(ctx, query, postId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		var comment entity.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.CreaterName,
			&comment.Avatar,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.Image,
		)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, comments, nil
}
