package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"

	"golang.org/x/crypto/bcrypt"
)

/*___________ THOS FUNCTIONS FOR AUTHENTICATION ___________*/

func (r *user) GetUserProfileById(ctx context.Context, targetId int) (entity.User, error) {
	var user entity.User
	requesterId := ctx.Value(entity.ContextID).(int)
	query := `
		SELECT 
			u.id, u.nickname, u.email, u.avatar, 
			u.first_name, u.last_name, u.birthday, u.about_me, 
			u.status,
			(SELECT COUNT(*) FROM follows WHERE followed_id = u.id) AS followers_count,
			(SELECT COUNT(*) FROM follows WHERE follower_id = u.id) AS following_count,
			CASE
				WHEN EXISTS (SELECT 1 FROM follows WHERE follower_id = $1 AND followed_id = $2) THEN 1
				WHEN EXISTS (SELECT 1 FROM notification WHERE sender_id = $1 AND receiver_id = $2 AND type = $3) THEN 2
				ELSE 0
			END AS following_status
		FROM users u
		WHERE u.id = $2

	`

	err := r.db.QueryRowContext(ctx, query, requesterId, targetId, entity.FollowingNotification).Scan(
		&user.ID, &user.Nickname, &user.Email, &user.Avatar,
		&user.First, &user.Last, &user.DateOfBirth, &user.AboutMe, &user.Status,
		&user.FollowersCount, &user.FollowingCount, &user.FollowingState)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}

	return user, nil
}

// this function is used to get user by username
func (r *user) GetUserByUsername(username string) (entity.User, error) {
	// SQL query that includes the counts and following state
	query := `SELECT * FROM users WHERE nickname = $1 OR email = $1`

	var user entity.User
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}
	err = stmt.QueryRow(username).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.First,
		&user.Last,
		&user.DateOfBirth,
		&user.Avatar,
		&user.Nickname,
		&user.AboutMe,
		&user.Status,
		&user.FollowersCount,
		&user.FollowingCount,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return user, err
		}
	}
	return user, nil
}

func (u *user) GetGroupById(ctx context.Context, groupId int) (entity.Group, error) {
	var group entity.Group
	query := `
		SELECT
			g.id, g.name
		FROM groups g
		WHERE g.id = $1
	`

	err := u.db.QueryRowContext(ctx, query, ctx.Value(entity.ContextID).(int), groupId).Scan(&group.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return group, fmt.Errorf("group not found")
		}
		return group, err
	}
	return group, nil
}

// this function is used to create new user
func (u *user) CreateUser(user entity.User) error {
	if user.Nickname.String != "" {
		user.Nickname.Valid = true
		ok, err := u.CheckUserByUsername(user.Nickname.String)
		if err != nil {
			return err
		}
		if ok {
			return config.ErrUserAlreadyExists
		}
	}
	query := `INSERT INTO users(
						Email,
						Password,
						first_name,
						last_name,
						birthday,
						Nickname,
						about_me,
						status,
						Avatar)
						VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.Email,
		user.Password,
		user.First,
		user.Last,
		user.DateOfBirth,
		user.Nickname,
		user.AboutMe,
		user.Status,
		user.Avatar)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) changeStatusRepo(ctx context.Context, id int) error {
	query := `UPDATE users SET status = 1 - status WHERE id = $1`
	_, err := u.db.ExecContext(ctx, query, id)
	return err
}

// this function is used to update user
func (r *user) UpdateUser(user entity.User) error {
	query := `UPDATE users SET
						Email = $1,
						Password = $2,
						first_name = $3,
						last_name = $4,
						birthday = $5,
						nick_name = $6,
						about_me = $7
						WHERE id = $8`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.Email,
		user.Password,
		user.First,
		user.Last,
		user.DateOfBirth,
		user.Nickname,
		user.AboutMe)
	if err != nil {
		return err
	}
	return nil
}

// this function is used to delete user by id
func (r *user) DeleteUser(id uint) error {
	query := `DELETE FROM users WHERE id = $1`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// this function is used to delete user by nickname
func (u *user) DeleteUserByNickName(nickName string) error {
	ok, err := u.CheckUserByUsername(nickName)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	query := `DELETE FROM users WHERE Nickname = $1`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(nickName)
	if err != nil {
		return err
	}
	return nil
}

/* ___________ THOSE FUNCS USED FOR CHECK USER CREDENTIALS ___________ */

// We'll use the Exists method to check if a user exists with a specific ID.
func (u *user) IsUserExist(id uint) (bool, error) {
	var exists bool
	stmt := "SELECT EXISTS(SELECT true FROM users WHERE id = ?)"
	err := u.db.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (u *user) authenticateRepo(email, password string) (int, error) {
	// u.loger.Info.Println("email:", email)
	// u.loger.Info.Println("password:", password)
	// Retrieve the id and hashed password associated with the given email. If
	// no matching email exists we return the ErrInvalidCredentials error.
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, password FROM users WHERE email = ? OR nickname  = ?"
	err := u.db.QueryRow(stmt, email, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, config.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	// Check whether the hashed password and plain-text password provided match.
	// If they don't, we return the ErrInvalidCredentials error.
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, config.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	// Otherwise, the password is correct. Return the user ID.
	return id, nil
}

// this function checks by email if the user exists
func (r *user) CheckUserByEmail(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exists bool
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return exists, err
	}
	err = stmt.QueryRow(email).Scan(&exists)
	if err != nil {
		return exists, err
	}
	return exists, nil
}

// this function checks by username if the user exists
func (r *user) CheckUserByUsername(username string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE Nickname = $1)`
	var exists bool
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return exists, err
	}
	err = stmt.QueryRow(username).Scan(&exists)
	if err != nil {
		return exists, fmt.Errorf("error executing query: %w", err)
	}
	return exists, nil
}

/*________________ THOSE FUNCS USED TO CHECK FOLOWING ____________ */
func (u *user) isFollowedBy(follower, followed int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = $1 AND followed_id = $2)`
	var exists bool
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return exists, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(follower, followed).Scan(&exists)
	if err != nil {
		return exists, fmt.Errorf("error executing query: %w", err)
	}

	return exists, nil
}

func (u *user) isFollowingEither(follower, followed int) (bool, error) {
	query := `SELECT EXISTS(
                SELECT 1 FROM follows WHERE (follower_id = $1 AND followed_id = $2) 
                OR (follower_id = $2 AND followed_id = $1)
              )`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	var exists bool
	err = stmt.QueryRow(follower, followed).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error executing query: %w", err)
	}

	return exists, nil
}

func (u *user) FollowRepository(followerId, followedId int) (bool, error) {
	// Check if the follow relationship already exists
	queryCheck := `SELECT 1 FROM follows WHERE follower_id = $1 AND followed_id = $2`
	var exists int
	err := u.db.QueryRow(queryCheck, followerId, followedId).Scan(&exists)

	if err == nil { // Row exists → Unfollow (delete)
		queryDelete := `DELETE FROM follows WHERE follower_id = $1 AND followed_id = $2`
		_, err = u.db.Exec(queryDelete, followerId, followedId)
		if err != nil {
			return false, fmt.Errorf("failed to unfollow: %w", err)
		}
		return false, nil
	} else if err != sql.ErrNoRows { // Any other error (DB issue)
		return false, fmt.Errorf("database error: %w", err)
	}

	// Row does not exist → Follow (insert)
	queryInsert := `INSERT INTO follows (follower_id, followed_id) VALUES ($1, $2)`
	_, err = u.db.Exec(queryInsert, followerId, followedId)
	if err != nil {
		return false, fmt.Errorf("failed to follow: %w", err)
	}
	return true, nil
}

func (u *user) CreateFollowNotification(ctx context.Context, senderId, receiverId int) (int, error) {
	query := `INSERT INTO notification (sender_id, receiver_id, type) 
	VALUES (?, ?, ?)`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, senderId, receiverId, entity.FollowingNotification)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// need some changes to follow up with the macro image
func (u *user) GroupContainsMember(groupId, userId int) (bool, error) {
	query := `SELECT 1 FROM members WHERE group_id = ? AND user_id = ? LIMIT 1`
	var exists int

	err := u.db.QueryRow(query, groupId, userId).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error checking group membership: %v", err)
	}

	return true, nil
}

/*___________ THOS FUNC USED FOR FOLLOWERS ___________*/
// this function is used to get followers by user id
func (u *user) GetFollowers(ctx context.Context, id int) ([]entity.User, error) {
	query := `
	SELECT users.id, first_name, last_name, status, nickname, avatar 
	FROM users
	INNER JOIN follows ON users.id = follows.follower_id
	WHERE follows.followed_id = ?`

	rows, err := u.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []entity.User
	for rows.Next() {
		var follower entity.User
		if err := rows.Scan(&follower.ID, &follower.First, &follower.Last, &follower.Status, &follower.Nickname, &follower.Avatar); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

// this function is used to get following by user id
func (u *user) GetFollowing(ctx context.Context, id int) ([]entity.User, error) {
	query := `
	SELECT users.id, first_name, last_name, status, nickname, avatar 
	FROM users
	INNER JOIN follows ON users.id = follows.followed_id 
	WHERE follows.follower_id = ?`

	rows, err := u.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []entity.User
	for rows.Next() {
		var follower entity.User
		if err := rows.Scan(&follower.ID, &follower.First, &follower.Last, &follower.Status, &follower.Nickname, &follower.Avatar); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

func (u *user) userNotificationRepo(ctx context.Context) ([]entity.Notification, int, error) {
	query := `SELECT 
                n.id,
                n.type,
                n.group_id,
                n.sender_id,
                n.receiver_id,
                n.event_id
              FROM notification n
              WHERE n.receiver_id = $1
              OR (
                  n.group_id IS NOT NULL 
                  AND EXISTS (
                      SELECT 1 FROM group_members gm
                      JOIN groups g ON gm.group_id = g.id
                      WHERE gm.member_id = $1
                      AND gm.group_id = n.group_id
                  )
              )
              ORDER BY n.id DESC`

	rows, err := u.db.QueryContext(ctx, query, ctx.Value(entity.ContextID).(int))
	if err != nil {
		return nil, http.StatusInternalServerError,
			fmt.Errorf("error querying notifications: %w", err)
	}
	defer rows.Close()

	var notifications []entity.Notification
	for rows.Next() {
		var (
			n entity.Notification
			groupId sql.NullInt64
			eventId sql.NullInt64
			receiverId sql.NullInt64
		)
		err = rows.Scan(
			&n.Id,
			&n.Type,
			&groupId,
			&n.SenderId,
			&receiverId,
			&eventId,
		)
		if err != nil {
			return nil, http.StatusInternalServerError,
				fmt.Errorf("error scanning notification: %w", err)
		}
		if groupId.Valid {
			n.GroupId = int(groupId.Int64)
		}
		if eventId.Valid {
			n.EventID = int(eventId.Int64)
		}
		if receiverId.Valid {
			n.ReceiverID = int(receiverId.Int64)
		}
		
		notifications = append(notifications, n)
	}

	if err = rows.Err(); err != nil {
		return nil, http.StatusInternalServerError,
			fmt.Errorf("rows iteration error: %w", err)
	}

	return notifications, http.StatusOK, nil
}

func (p *user) GetUserPostsRep(ctx context.Context, cursor entity.Cursor) ([]entity.Post, error) {
	query := `
		SELECT
			post.id,
			user.avatar,
			post.content,
			post.image,
			(SELECT nickname FROM users AS u WHERE post.user_id=u.id) AS creator,
			post.created_at
		FROM posts AS post
		INNER JOIN users AS user ON user.id = post.user_id
		WHERE user.id = $1
		AND (
			$2 IS NULL
			OR (post.created_at < $2 AND post.id < $3)
			OR (post.created_at = $2 AND post.id  < $3)
		)
		ORDER BY post.created_at DESC, post.id DESC
		LIMIT $4;
	`

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows *sql.Rows
	if cursor.Time == nil || cursor.LastId == nil {
		// First page — pass NULL cursor
		rows, err = stmt.QueryContext(ctx, cursor.ID, nil, nil, cursor.Limit)
	} else {
		rows, err = stmt.QueryContext(ctx, cursor.ID, cursor.Time, cursor.LastId, cursor.Limit)
	}
	if err != nil {
		return nil, err
	}

	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(&post.ID, &post.Avatar, &post.Content, &post.Image, &post.UserName, &post.CreatedAt)
		if err != nil {
			p.loger.Error.Println("error occur while scan post info", err)
			continue
		}
		posts = append(posts, post)
	}

	return posts, nil
}
