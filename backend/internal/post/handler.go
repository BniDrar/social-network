package post

import (
	"database/sql"
	"net/http"

	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type Post interface {
	Post(w http.ResponseWriter, r *http.Request)
	GetPosts(w http.ResponseWriter, r *http.Request)
	React(w http.ResponseWriter, r *http.Request)
}

type post struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

func Newpost(dep *config.Dependencies) Post {
	return &post{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
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
