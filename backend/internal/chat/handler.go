package chat

import (
	"database/sql"
	"net/http"
	"socialNetwork/pkg/websocket"
)

type chat struct {
	serv *Service
}

type Chat interface {
	WebSocket(w http.ResponseWriter, r *http.Request)
}

func NewChat(db *sql.DB, hub *websocket.Hub) Chat {
	repo:= NewRepo(db)
	serv:= NewService(repo, hub)
	return &chat{serv: serv}
}

func (c *chat) WebSocket(w http.ResponseWriter, r *http.Request) {
	//upgrade
	
	//
}