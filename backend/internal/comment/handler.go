package comment

import (
	"database/sql"
	"net/http"
	"socialNetwork/pkg/websocket"
)

type comment struct {
	serv *Service
}

type Comment interface {
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Vote(w http.ResponseWriter, r *http.Request)
}

func NewComment(db *sql.DB, hub *websocket.Hub) Comment {
	repo:= NewRepo(db)
	serv:= NewService(repo, hub)
	return &comment{serv: serv}
}

func (c *comment) Post(w http.ResponseWriter, r *http.Request) {}

func (c *comment) Get(w http.ResponseWriter, r *http.Request) {}

func (c *comment) Vote(w http.ResponseWriter, r *http.Request) {}
