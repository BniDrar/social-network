package user

import (
	"errors"
	"log"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/utils"
)

func (s *user) LoginService(user entity.Credentials) (string, error, int) {
	//  check credentials
	if err := utils.ValidateLoginCredentials(user); err != nil {
		return "", err, http.StatusBadRequest
	}
	// get user from db
	u, err := s.GetUserByUsername(user.Username)
	if err != nil {
		return "", errors.New("invalid username or password"), http.StatusBadRequest
	}
	// compare password
	utils.ComparePasswords(u.Password, user.Password)
	// generate token
	token, err := utils.GenerateToken()
	if err != nil {
		return "", errors.New("internal server error"), http.StatusInternalServerError
	}
	return token, nil, http.StatusOK
}

func (s *user) RegisterService(user entity.User) (error, int) {
	// check credentials
	if err := utils.ValidateRegisterCredentials(user); err != nil {
		return err, http.StatusBadRequest
	}
	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	user.Password = hashedPassword
	// chek if user Email already exists
	_, err = s.GetUserByUsername(user.Email)
	if err == nil {
		return errors.New("email already exists"), http.StatusBadRequest
	}
	// check if user Nickname already exists
	_, err = s.GetUserByUsername(user.Nickname)
	if err == nil {
		return errors.New("nickname already exists"), http.StatusBadRequest
	}
	// save user to db
	err = s.CreateUser(user)
	if err != nil {
		log.Println("hre %v ", err)
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *user) LogoutService(user entity.User) error {
	// do something
	return nil
}

func (s *user) ProfileService(user entity.User) error {
	// do something
	return nil
}

func (s *user) FollowService(user entity.User) error {
	// do something
	return nil
}

func (s *user) FollowersService(user entity.User) error {
	// do something
	return nil
}

func (s *user) IsExistsService(user uint) bool {
	// do something
	return false
}
