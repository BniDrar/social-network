package utils

import (
	"fmt"
	"net/mail"
	"regexp"
	"unicode"

	"socialNetwork/entity"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// this function is used to validate the user credentials before saving to the database
func ValidateLoginCredentials(user entity.Credentials) error {
	if err := ValidateUsername(user.Username); err != nil {
		return err
	}
	if !isValidPassWord(user.Password) {
		return fmt.Errorf("invalid password")
	}
	return nil
}

//  this function the register credentials before saving to the database
func ValidateRegisterCredentials(user entity.User) error {
	if user.First == "" || user.Last == "" {
		return fmt.Errorf("invalid first or last name")
	}
	if !isValidNickName(user.Nickname) {
		return fmt.Errorf("invalid nickname")
	}
	if ok, err := isValidateEmail(user.Email); !ok {
		return fmt.Errorf("invalid email %v", err)
	}
	if !isValidPassWord(user.Password) {
		return fmt.Errorf("invalid password")
	}
	return nil
}

// this if know if the user enter email or nickname in the username field and validate it
func ValidateUsername(username string) error {
	if ok, err := isValidateEmail(username); ok {
		return fmt.Errorf("invalid email%v", err)
	}
	if !isValidNickName(username) {
		return fmt.Errorf("invalid nickname")
	}
	return nil
}

// this function is used to validate the email
func isValidateEmail(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, err
	}
	return true, nil
}

// this function is used to validate the password
func isValidPassWord(password string) bool {
	if len(password) < 8 {
		return false
	}
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsNumber(r):
			hasNumber = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}
	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return false
	}
	return true
}

// this function is used to validate the username
func isValidNickName(name string) bool {
	if len(name) < 3 {
		return false
	}
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", name); !ok {
		return false
	}
	return true
}

// this function is used to compare the hashed password with the plain password
func ComparePasswords(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password %v", err)
	}
	return nil
}

// this function is used to hash the password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error while hashing the password %v", err)
	}
	return string(hash), nil
}

// this function is used to generate the token
// this still need to be implemented in the future
func GenerateToken() (string, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("error while generating the token %v", err)
	}
	return token.String(), nil
}
