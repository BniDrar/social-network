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
			OR (post.user_id = $1);
		ORDER BY post.created_at DESC`
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

func (p *post) getAllPostsRepo(ctx context.Context, userID int, cursor entity.Cursor) ([]entity.Post, int, error) {
	query := `
		SELECT
			post.id,
			user.avatar,
			user.first_name,
			user.last_name,
			post.content,
			post.image,
			(SELECT COUNT(*) FROM engagement AS eng WHERE eng.user_id = $1 AND eng.post_id = post.id) AS likes_count,
			(SELECT COUNT(*) FROM comments AS c WHERE c.post_id = post.id) AS comments,
			(SELECT nickname FROM users AS u WHERE post.user_id = u.id) AS creator,
			CASE 
				WHEN EXISTS (
					SELECT 1 FROM engagement eng 
					WHERE eng.user_id = $1 AND eng.post_id = post.id
				) THEN 1 ELSE 0
			END AS engaged,
			post.created_at

		FROM posts AS post
		INNER JOIN users AS user ON user.id = post.user_id
		LEFT JOIN groups AS "group" ON "group".id = post.group_id
		LEFT JOIN group_members AS gm ON gm.member_id = $1 AND gm.group_id = "group".id
		LEFT JOIN follows AS follow 
			ON (follow.followed_id = post.user_id AND 
				follow.follower_id = $1 AND 
				post.status = 0 AND 
				post.group_id IS NULL)
		WHERE
			(
				(post.status = 0 AND gm.member_id IS NOT NULL)
				OR (post.status = 1 AND follow.follower_id = $1)
				OR (post.status = 2)
				OR (post.user_id = $1)
			)
			AND (
				$2 IS NULL
				OR (post.created_at < $2 AND post.id < $3)
				OR (post.created_at = $2 AND post.id  < $3)
			)
		ORDER BY post.created_at DESC, post.id DESC
		LIMIT $4;
	`

	prep, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer prep.Close()
	if (cursor.Time == nil && cursor.LastId != nil) || (cursor.Time != nil && cursor.LastId == nil) {
		return nil, http.StatusBadRequest, errors.New("invalid cursor: both time and id must be set together")
	}

	var rows *sql.Rows
	if cursor.Time == nil || cursor.LastId == nil {
		// First page → pass NULLs
		rows, err = prep.QueryContext(ctx, userID, nil, nil, cursor.Limit)
	} else {
		rows, err = prep.QueryContext(ctx, userID, cursor.Time, cursor.LastId, cursor.Limit)
	}
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid query")
	}

	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(
			&post.ID, &post.Avatar, &post.First, &post.Last,
			&post.Content, &post.Image, &post.LikesCount, &post.Comments,
			&post.UserName, &post.Engagement, &post.CreatedAt, // <-- Don't forget created_at!
		)
		if err != nil {
			p.loger.Error.Println("error occur while scan post info", err)
			continue
		}
		posts = append(posts, post)
	}

	return posts, http.StatusOK, nil
}

func (p *post) getPostsByGroupId(ctx context.Context, userId int, cursor entity.Cursor) ([]entity.Post, error) {
	if cursor.ID <= 0 {
		return nil, errors.New("group ID is required")
	}

	query := `
		SELECT
			post.id,
			user.avatar,
			user.first_name,
			user.last_name,
			post.content,
			post.image,
			(SELECT COUNT(*) FROM engagement AS eng WHERE eng.user_id = $1 AND eng.post_id = post.id) AS likes_count,
			(SELECT COUNT(*) FROM comments AS c WHERE c.post_id = post.id) AS comments,
			(SELECT nickname FROM users AS u WHERE post.user_id = u.id) AS creator,
			CASE 
				WHEN EXISTS (
					SELECT 1 FROM engagement eng 
					WHERE eng.user_id = $1 AND eng.post_id = post.id
				) THEN 1 ELSE 0
			END AS engaged,
			post.created_at

		FROM posts AS post
		INNER JOIN users AS user ON user.id = post.user_id
		LEFT JOIN group_members AS gm ON gm.member_id = $1 AND gm.group_id = post.group_id

		WHERE
			post.group_id = $2
			AND gm.member_id IS NOT NULL
			AND (
				$3 IS NULL
				OR (post.created_at < $3 AND post.id < $4)
				OR (post.created_at = $3 AND post.id < $4)
			)
		ORDER BY post.created_at DESC, post.id DESC
		LIMIT $5;
	`

	prep, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer prep.Close()

	// Validate cursor
	if (cursor.Time == nil && cursor.LastId != nil) || (cursor.Time != nil && cursor.LastId == nil) {
		return nil, errors.New("invalid cursor: both time and id must be set together")
	}

	var rows *sql.Rows
	if cursor.Time == nil || cursor.LastId == nil {
		// First page → pass NULLs
		rows, err = prep.QueryContext(ctx, userId, cursor.ID, nil, nil, cursor.Limit)
	} else { 
		rows, err = prep.QueryContext(ctx, userId, cursor.ID, cursor.Time, cursor.LastId, cursor.Limit)
	}
	if err != nil {
		return nil, err
	}

	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(
			&post.ID, &post.Avatar, &post.First, &post.Last,
			&post.Content, &post.Image, &post.LikesCount, &post.Comments,
			&post.UserName, &post.Engagement, &post.CreatedAt,
		)
		if err != nil {
			p.loger.Error.Println("error occur while scan post info", err)
			continue
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *post) GetPostRepo(ctx context.Context, post_id int) (entity.Post, int, error) {
	user_id := ctx.Value(entity.ContextID).(int)
	prep, err := p.db.PrepareContext(ctx, `SELECT
		post.id,
		user.avatar,
		post.content,
		post.image,
		user.nickname AS creator,
		(SELECT COUNT(*) FROM engagement AS eng WHERE eng.post_id = post.id) AS likes_count,
		(SELECT COUNT(*) FROM comments AS c WHERE c.post_id = post.id) AS comments,
		CASE 
			WHEN EXISTS (
				SELECT 1 FROM engagement AS eng 
				WHERE eng.user_id = $1 AND eng.post_id = post.id
			) THEN 1 ELSE 0
		END AS engaged
	FROM posts AS post 
	INNER JOIN users AS user ON user.id = post.user_id
	WHERE (post.id = $2);
`)
	if err != nil {
		return entity.Post{}, http.StatusInternalServerError, err
	}
	var post entity.Post
	res := prep.QueryRowContext(ctx, user_id, post_id)
	err = res.Scan(&post.ID, &post.Avatar, &post.Content, &post.Image, &post.UserName, &post.LikesCount, &post.Comments, &post.Engagement)
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
		r.loger.Error.Println(err)
		return 0, err
	}

	// Add viewers to the group
	for _, viewerID := range viewers {
		_, err := tx.Exec(`INSERT INTO group_members (group_id, member_id) VALUES ($1, $2)`, groupID, viewerID)
		if err != nil {
			return 0, errors.New("error while executing the transaction: " + err.Error())
		}
	}

	return groupID, nil
}

func (p *post) PostEngagementRepo(ctx context.Context, userId, postId int) error {
	countQuery := `SELECT COUNT(*) FROM engagement WHERE user_id = $1 AND post_id = $2`
	insertQuery := `INSERT INTO engagement (user_id, post_id, status) VALUES ($1, $2, 1)`
	deleteQuery := `DELETE FROM engagement WHERE user_id = $1 AND post_id = $2`

	// Prepare count statement
	countStmt, err := p.db.PrepareContext(ctx, countQuery)
	if err != nil {
		return fmt.Errorf("failed to prepare count query: %w", err)
	}
	defer countStmt.Close()

	var count int
	err = countStmt.QueryRowContext(ctx, userId, postId).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to count engagement: %w", err)
	}

	var stmt *sql.Stmt

	if count > 0 {
		stmt, err = p.db.PrepareContext(ctx, deleteQuery)
		if err != nil {
			return fmt.Errorf("failed to prepare delete query: %w", err)
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, userId, postId)
		if err != nil {
			return fmt.Errorf("failed to delete engagement: %w", err)
		}
	} else {
		stmt, err = p.db.PrepareContext(ctx, insertQuery)
		if err != nil {
			return fmt.Errorf("failed to prepare insert query: %w", err)
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, userId, postId)
		if err != nil {
			return fmt.Errorf("failed to insert engagement: %w", err)
		}
	}

	return nil
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
