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
	// upgrade
	log.Println("WebSocket endpoint hit")
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins (use cautiously in production)
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.loger.Error.Println("Error while upgrading connection:", err)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Error while upgrading connection"})
		return
	}
	defer conn.Close()
	log.Println("Client connected")
	// get user id from session
	userId := r.Context().Value("userID")
	if userId == nil {
		c.loger.Error.Println("User ID not found in session")
		//w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//
	WsListing(conn, r.Context())
	c.loger.Info.Println("Client disconnected")
}
