package user

import (
	"database/sql"
	"fmt"
	"log"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
)

/*___________ THOS FUNCTIONS FOR AUTHENTICATION ___________*/

// this function is used to get user by username
func (r *user) GetUserByUsername(username string) (entity.User, error) {
	// SQL query that includes the counts and following state
	query := `SELECT * FROM users WHERE nickname = $1 OR email = $1`

	var user entity.User
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("err 21", err)
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

// this function is used to create new user
func (u *user) CreateUser(user entity.User) error {
	// ok, err := u.IsExistsService(int(user.ID))
	ok, err := u.CheckUserByUsername(user.Nickname)
	if err != nil {
		return err
	}
	if ok {
		return config.ErrUserAlreadyExists
	}
	query := `INSERT INTO users(
						Email,
						Password,
						first_name,
						last_name,
						birthday,
						Nickname,
						about_me,
						Avatar)
						VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
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
		user.Avatar)
	if err != nil {
		return err
	}
	return nil
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

// this function is used to delete user by id
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

/* ___________ THOS FUNC USED FOR CHECK USER CREDENTIALS ___________ */

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

func (u *user) isFollowedBy(follower, followed int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM follows WHERE folowwer_id = $1 AND followed_id = $2)`
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

/* ___________ THOS FUNC USED FOR FOLLOWERS ___________ */

// this function is used to get followers by user id
// func (r *user) GetFollowers(id uint) ([]entity.User, error) {}
// this function is used to get following by user id
// func (r *user) GetFollowing(id uint) ([]entity.User, error) {}

// this function is used to follow user
// func (r *user) Follow(follower, following uint) error {}
// this function is used to unfollow user
// func (r *user) Unfollow(follower, following uint) error {}
// this function is used to get followers count
// func (r *user) GetFollowersCount(id uint) (uint, error) {}
// this function is used to get following count
// func (r *user) GetFollowingCount(id uint) (uint, error) {}
