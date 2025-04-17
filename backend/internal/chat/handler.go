package chat

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	scs "socialNetwork/pkg/sessions"
	ws "socialNetwork/pkg/websocket"

	"github.com/gorilla/websocket"
)

type chat struct {
	Hub            *ws.Hub
	db             *sql.DB
	loger          loger.CstmLogger
	sessionManager *scs.SessionManager
}

type Chat interface {
	WebSocket(w http.ResponseWriter, r *http.Request)
}

func NewChat(dep *config.Dependencies) Chat {
	return &chat{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (use cautiously in production)
	},
}

func (c *chat) WebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket endpoint hit")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.loger.Error.Println("Error while upgrading connection:", err)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Error while upgrading connection"})
		return
	}
	defer conn.Close()
	log.Println("Client connected")
	// get user id from session
	userId := r.Context().Value(entity.ContextID).(int)
	c.loger.Info.Println("user Id", userId)
	c.loger.Info.Println("user Ip", conn.RemoteAddr().String())
	WsListing(conn, r.Context())
	c.loger.Info.Println("Client disconnected")
}
