package post

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"socialNetwork/entity"
)

func (p *post) GroupMember(ctx context.Context, userID, groupID int) bool {
	query := `SELECT
		    user.id
		FROM
			users AS user
		LEFT JOIN group_members AS gm ON gm.group_id = $2
			AND gm.member_id = $1
		WHERE
		    ($2 IS NOT NULL AND gm.member_id IS NOT NULL)`
	smtp, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return false
	}
	row := smtp.QueryRowContext(ctx, userID, groupID)
	var res int
	err = row.Scan(&res)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *post) Repo_UserCanPost(ctx context.Context, id, postid int) bool {
	query := `SELECT
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
			OR (post.user_id = $1);`
	smtp, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return false
	}
	row := smtp.QueryRowContext(ctx, id, postid)
	var res int
	err = row.Scan(&res)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *post) Repo_GetAll(ctx context.Context, id int) (posts []entity.Post, err error) {
	prep, err := p.db.PrepareContext(ctx, `SELECT
			post.id,
		    user.avatar,
		    post.content,
			post.image,
		    (SELECT nickname FROM users AS u WHERE post.user_id=u.id) AS creator
		FROM
		    posts AS post
		INNER JOIN users AS user ON user.id= post.user_id
		LEFT JOIN groups AS "group" ON "group".id = post.group_id
		LEFT JOIN group_members AS gm ON gm.member_id = $1 AND gm.group_id = "group".id
		LEFT JOIN follows AS follow 
		    ON (follow.followed_id = post.user_id AND 
		        follow.follower_id = $1 AND 
		        post.status = 0 AND 
		        post.group_id IS NULL)
		WHERE
		    (post.status = 0 AND gm.member_id IS NOT NULL )
		    OR (post.status = 1 AND follow.follower_id = $1 )
			OR (post.status = 2 )
			OR (post.user_id = $1);`)
	if err != nil {
		return
	}
	res, err := prep.QueryContext(ctx, id)
	if err != nil {
		fmt.Println("erro here", err)
		return
	}
	for res.Next() {
		fmt.Println("Gitted")
		post := entity.Post{}
		err := res.Scan(&post.ID, &post.Avatar, &post.Content, &post.Image, &post.UserName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(post)
		posts = append(posts, post)
	}
	return
}

func (p *post) GetPostRepo(ctx context.Context, post_id int) (entity.Post, int, error) {
	prep, err := p.db.PrepareContext(ctx, `SELECT
		    post.id,
		    user.avatar,
		    post.content,
			post.image,
		    user.nickname AS creator
		FROM posts AS post 
		INNER JOIN users AS user ON user.id = post.user_id
		WHERE (post.id=$1);`)
	if err != nil {
		return entity.Post{}, http.StatusInternalServerError, err
	}
	var post entity.Post
	res := prep.QueryRowContext(ctx, post_id)
	err = res.Scan(&post.ID, &post.Avatar, &post.Content, &post.Image, &post.UserName)
	if err != nil {
		return entity.Post{}, http.StatusInternalServerError, err
	}
	return post, http.StatusOK, nil
}

func (r *post) SavePost(ctx context.Context, userID int, post entity.Post) (int, int, error) {
	// Start a transaction for creating post and group
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	defer tx.Rollback() // Rollback if there is any error

	// Insert post first
	var postID int
	err = tx.QueryRow(`INSERT INTO posts (user_id, content, image, group_id, status) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		userID, post.Content, post.Image, post.GroupID, post.Status).Scan(&postID)
	if err != nil {
		return 0, http.StatusBadRequest, errors.New("error in query row: " + err.Error())
	}

	// Handle custom group creation if status is 'custom'
	if post.Status == entity.PostStatusCustom {
		groupID, err := r.createCustomGroup(tx, post.AllowedViewers)
		if err != nil {
			return 0, http.StatusBadRequest, err
		}

		// Update post with the new group ID
		_, err = tx.Exec(`UPDATE posts SET group_id = $1 WHERE id = $2`, groupID, postID)
		if err != nil {
			return 0, http.StatusInternalServerError, err
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return 0, http.StatusInternalServerError, errors.New("error while commiting the transaction" + err.Error())
	}

	return postID, http.StatusCreated, nil
}

func (r *post) createCustomGroup(tx *sql.Tx, viewers []int) (int, error) {
	// Create new custom group
	var groupID int
	err := tx.QueryRow(`INSERT INTO groups (type) VALUES (?) RETURNING id`, entity.PostStatusCustom).Scan(&groupID)
	if err != nil {
		return 0, err
	}

	// Add viewers to the group
	for _, viewerID := range viewers {
		_, err := tx.Exec(`INSERT INTO group_members (group_id, member_id) VALUES ($1, $2)`, groupID, viewerID)
		if err != nil {
			return 0, errors.New("error while executing the transaction: "+ err.Error())
		}
	}

	return groupID, nil
}

func (p *post) Repo_React(ctx context.Context, user_id int, react entity.Vote) (err error) {
	prep, err := p.db.PrepareContext(ctx, `INSERT INTO engagements
			(user_id, post_id, status)
		VALUES
			($1, $2, $3)`)
	if err != nil {
		return
	}
	_, err = prep.ExecContext(ctx, user_id, react.ID, react.Status)
	return
}

func (p *post) GetPostsByUserID(ctx context.Context, user_id int, username string) (posts []entity.Post, err error) {
	prep, err := p.db.PrepareContext(ctx, `SELECT
			post.id,
		    user.avatar,
		    post.content,
			post.image,
		    (SELECT nickname FROM users AS u WHERE post.user_id=u.id) AS creator,
			post.status,
			"group".name
		FROM
		    posts AS post
		INNER JOIN users AS user ON user.id= post.user_id
		LEFT JOIN groups AS "group" ON "group".id = post.group_id
		LEFT JOIN group_members AS gm ON gm.member_id = $1 AND gm.group_id = "group".id
		LEFT JOIN follows AS follow 
		    ON (follow.followed_id = post.user_id AND 
		        follow.follower_id = $1 AND 
		        post.status = 0 AND 
		        post.group_id IS NULL)
		WHERE
		    (post.status = 0 AND gm.member_id IS NOT NULL AND post.id = $2 AND user.nickname = $2)
		    OR (post.status = 1 AND follow.follower_id = $1 AND post.id = $2 AND user.nickname = $2)
			OR (post.status = 2 AND post.id = $2 AND user.nickname = $2)
			OR (post.user_id = $1 AND user.nickname = $2)`)
	if err != nil {
		return
	}
	res, err := prep.QueryContext(ctx, user_id, username)
	if err != nil {
		return
	}
	for res.Next() {
		post := entity.Post{}
		err := res.Scan(&post.ID, &post.Avatar, &post.Content, &post.Image, &post.UserName, &post.Status, &post.GroupName)
		if err != nil {
			fmt.Println(entity.WhereIsError() + " " + err.Error())
			continue
		}
		posts = append(posts, post)
	}
	return
}
