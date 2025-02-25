package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"socialNetwork/entity"
)

type User interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	Follow(w http.ResponseWriter, r *http.Request)
	Followers(w http.ResponseWriter, r *http.Request)
}

type user struct {
	serv *Service
}

func NewUser(db *sql.DB) User {
	repo := NewRepo(db)
	serv := NewService(repo)
	return &user{serv: serv}
}

func (u *user) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get user data from request
	var User entity.Credentials
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		http.Error(w, "Error while parsing request", http.StatusBadRequest)
		return
	}
	// call service
	token, err, status := u.serv.Login(User)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	// send response and set token in cookie
	http.SetCookie(w, &http.Cookie{
		Name: token,
		Value: token,
	})
  w.WriteHeader(http.StatusOK)
}

func (u *user) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	User := entity.User{}
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		http.Error(w, "Error while parsing request", http.StatusBadRequest)
		return
	}
	err,status := u.serv.Register(User)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(status)
}

func (u *user) Logout(w http.ResponseWriter, r *http.Request) {}

func (u *user) Profile(w http.ResponseWriter, r *http.Request) {}

func (u *user) Follow(w http.ResponseWriter, r *http.Request) {}

func (u *user) Followers(w http.ResponseWriter, r *http.Request) {}
