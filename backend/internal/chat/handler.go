package chat

import (
	"database/sql"
	"fmt"
	"net/http"

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

	wupgrade := w
	if u, ok := w.(interface{ Unwrap() http.ResponseWriter }); ok {
		wupgrade = u.Unwrap()
	}

	c.loger.Info.Printf("w's type is %T\n", w)

	ws, err := upgrader.Upgrade(wupgrade, r, nil)
	if err != nil {
		c.loger.Error.Println(err)
		return
	}

	c.loger.Info.Printf("Web client connected from %s", r.RemoteAddr)

	fmt.Println("done", ws)
}
