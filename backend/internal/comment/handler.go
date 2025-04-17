package comment

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type comment struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

type Comment interface {
	AddComment(w http.ResponseWriter, r *http.Request)
	GetComments(w http.ResponseWriter, r *http.Request)
	VoteComment(w http.ResponseWriter, r *http.Request)
}

/*               app     */
func NewComment(dep *config.Dependencies) Comment {
	return &comment{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

func (c *comment) AddComment(w http.ResponseWriter, r *http.Request) {}

func (c *comment) GetComments(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		UserID := r.Context().Value(entity.ContextID).(int)
		var Comment entity.Comment
		err := json.NewDecoder(r.Body).Decode(&Comment)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
			return
		}
		Comment.UserID = UserID
		// ---------------------------------- Service Add Comment ---------------------------------------

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Method Not Allowed"})
		return
	}
}

func (c *comment) VoteComment(w http.ResponseWriter, r *http.Request) {}
