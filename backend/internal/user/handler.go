package user

import (
	"database/sql"
	"fmt"
	"net/http"
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
	repo:= NewRepo(db)
	serv:= NewService(repo)
	return &user{serv: serv}
}

func (u *user) Login(w http.ResponseWriter, r *http.Request) {}

func (u *user) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "something")
}

func (u *user) Logout(w http.ResponseWriter, r *http.Request) {}

func (u *user) Profile(w http.ResponseWriter, r *http.Request) {}

func (u *user) Follow(w http.ResponseWriter, r *http.Request) {}

func (u *user) Followers(w http.ResponseWriter, r *http.Request) {}