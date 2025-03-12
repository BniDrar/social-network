package user

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (s *user) LoginService(user entity.Credentials) (string, int, error) {
	//  check credentials
	if err := utils.ValidateLoginCredentials(user); err != nil {
		return "", http.StatusBadRequest, err
	}
	// get user from db
	u, err := s.GetUserByUsername(user.Username)
	if err != nil {
		return "", http.StatusBadRequest, errors.New("invalid username or password")
	}
	// compare password
	utils.ComparePasswords(u.Password, user.Password)
	// generate token
	token, err := utils.GenerateToken()
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("internal server error")
	}
	return token, http.StatusOK, nil
}

func (s *user) RegisterService(user entity.User) (int, error) {
	// check credentials
	if err := utils.ValidateRegisterCredentials(user); err != nil {
		return http.StatusBadRequest, err
	}
	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	user.Password = hashedPassword
	// chek if user Email already exists
	_, err = s.GetUserByUsername(user.Email)
	if err != nil {
		return http.StatusBadRequest, errors.New("email already exists")
	}
	// check if user Nickname already exists
	_, err = s.GetUserByUsername(user.Nickname)
	if err != nil {
		return http.StatusBadRequest, errors.New("nickname already exists")
	}
	// save user to db
	err = s.CreateUser(user)
	if err != nil {
		if errors.Is(err, config.ErrUserAlreadyExists) {
			return http.StatusBadRequest, err
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (u *user) Authenticate(email, password string) (int, error) {
	// Retrieve the id and hashed password associated with the given email. If
	// no matching email exists we return the ErrInvalidCredentials error.
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, hashed_password FROM users WHERE email = ? OR username = ?"
	err := u.db.QueryRow(stmt, email).Scan(&id, &hashedPassword)
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

// We'll use the Exists method to check if a user exists with a specific ID.
func (u *user) IsExistsService(id int) (bool, error) {
	var exists bool
	stmt := "SELECT EXISTS(SELECT true FROM users WHERE id = ?)"
	err := u.db.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}

func (u *user) LogoutService(user entity.User) error {
	// do something
	return nil
}

func (u *user) UserProfile(ctx context.Context, nickname string) (int, user, error) {
	u.CheckUserByUsername(nickname)
	return 0, user{}, nil
}

func (u *user) FollowService(user entity.User) error {
	// do something
	return nil
}

func (u *user) FollowersService(user entity.User) error {
	// do something
	return nil
}

// func (s *user) FollowersService(user entity.User) error {
// 	// do something
// 	return nil
// }

// // func (s *user) IsExistsService(user uint) bool {
// // 	// do something
// // 	return false
// // }

// func (s *user) DeleteUserService(user entity.User) error {

// 	return nil
// }
