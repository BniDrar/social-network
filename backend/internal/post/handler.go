package post

import (
	"database/sql"
	"net/http"
)

type Post interface {
	Post(w http.ResponseWriter, r *http.Request)
	GetPosts(w http.ResponseWriter, r *http.Request)
	React(w http.ResponseWriter, r *http.Request)
}

type post struct {
	serv *Service
}

func Newpost(db *sql.DB) Post {
	repo:= NewRepo(db)
	serv:= NewService(repo)
	return &post{serv: serv}
}


func (u *post) GetPosts(w http.ResponseWriter, r *http.Request) {}

func (u *post) React(w http.ResponseWriter, r *http.Request) {}

func (u *post) Post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "Post":
		u.CreatePost(w, r)
	case "Get":
		u.GetPost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (u *post) CreatePost(w http.ResponseWriter, r *http.Request) {}

func (u *post) GetPost(w http.ResponseWriter, r *http.Request) {}
