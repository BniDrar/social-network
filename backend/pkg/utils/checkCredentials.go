package utils

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
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

// this function the register credentials before saving to the database
func ValidateRegisterCredentials(user entity.User) error {
	if !IsValidName(user.First) || !IsValidName(user.Last) {
		return fmt.Errorf("invalid first or last name")
	}
	// if !isValidNickName(user.Nickname) {
	// 	return fmt.Errorf("invalid nickname")
	// }
	if !isValidateEmail(user.Email) {
		return fmt.Errorf("invalid email %v", user.Email)
	}
	if !isValidPassWord(user.Password) {
		return fmt.Errorf("invalid password")
	}
	return nil
}

// this if know if the user enter email or nickname in the username field and validate it
func ValidateUsername(username string) error {
	// Check if input looks like an email (contains @)
	if strings.Contains(username, "@") {
		if isValidateEmail(username) {
			return nil
		}
		return fmt.Errorf("invalid email format: %s (example: user@domain.com)", username)
	}

	// If no @, treat as nickname
	if isValidNickName(username) {
		return nil
	}
	return fmt.Errorf("invalid nickname format: %s (must be 4-20 alphanumeric characters)", username)
}

// this function is used to validate the email
func isValidateEmail(email string) bool {
	m, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	if m.Address != email {
		return false
	}
	return true
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
	if name == "" || len(name) < 5 || len(name) > 20 {
		return false
	}
	// Compile regex once at package level for better performance
	var nickNameRegex = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9.]*[a-zA-Z0-9]$`)
	// Early exit if pattern doesn't match
	if !nickNameRegex.MatchString(name) {
		return false
	}

	// Check for consecutive periods
	return !strings.Contains(name, "..")
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
