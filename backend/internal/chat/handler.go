package chat

import (
	"database/sql"
	"encoding/json"
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
	GetUserContacts(w http.ResponseWriter, r *http.Request)
	GetOnlineUsers(w http.ResponseWriter, r *http.Request)
	GetMessages(w http.ResponseWriter, r *http.Request)
}

func NewChat(dep *config.Dependencies) Chat {
	return &chat{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (use cautiously in production)
	},
}

func (c *chat) GetOnlineUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	contacts, status, err := c.getOnlineUsers(r.Context())
	w.WriteHeader(status)
	if err != nil {
		c.loger.Error.Println("error while getting contacts: ", err)
		return
	}
	json.NewEncoder(w).Encode(&contacts)
}

func (c *chat) GetUserContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	contacts, status, err := c.getUserContactsService(r.Context())
	w.WriteHeader(status)
	if err != nil {
		c.loger.Error.Println("error while getting contacts: ", err)
		return
	}
	json.NewEncoder(w).Encode(&contacts)
}

func (c *chat) GetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var Cursor entity.Cursor
	err := json.NewDecoder(r.Body).Decode(&Cursor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}

	status, messages, err := c.getMessagesService(r.Context(), Cursor)
	w.WriteHeader(status)
	if err != nil {
		c.loger.Error.Println("error getting messages: ", err)
		return
	}
	json.NewEncoder(w).Encode(&messages)
}

func (c *chat) WebSocket(w http.ResponseWriter, r *http.Request) {
	// upgrade
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
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

	//get user id from session
	userId := r.Context().Value(entity.ContextID)
	if userId == nil {
		c.loger.Error.Println("User ID not found in session")
		//w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c.loger.Info.Printf("Client %d connected\n", userId)
	c.WsListing(conn, r.Context())
	c.loger.Info.Printf("Client %d disconnected\n", userId)
}
