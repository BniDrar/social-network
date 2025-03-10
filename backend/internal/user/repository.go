package user

import (
	"socialNetwork/entity"
)

/*___________ THOS FUNCTIONS FOR AUTHENTICATION ___________*/

// this function is used to get user by username
func (r *user) GetUserByUsername(username string) (entity.User, error) {
	query := `SELECT * FROM user WHERE Nickname = $1 OR Email = $1`
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
		&user.Status)
	if err != nil {
		return user, err
	}
	return user, nil
}

// this function is used to create new user
func (r *user) CreateUser(user entity.User) error {
	query := `INSERT INTO user(
						Email,
						Password,
						First,
						Last,
						Date_Of_Birth,
						Nickname,
						About_Me,
						Avatar)
						VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
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
						First = $3,
						Last = $4,
						Date_Of_Birth = $5,
						Nickname = $6,
						About_Me = $7
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
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	var exists bool
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return exists, err
	}
	err = stmt.QueryRow(username).Scan(&exists)
	if err != nil {
		return exists, err
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
