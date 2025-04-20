package comment

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/utils"
	"socialNetwork/pkg/websocket"

	"github.com/google/uuid"
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

func (c *comment) AddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var commnt entity.Comment
	err := json.NewDecoder(r.Body).Decode(&commnt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}

	fileName := uuid.NewString()
	file, fileHeader, err := r.FormFile("image")
	err = utils.FileUpload(fileName, file, fileHeader, err)
	if err == nil {
		commnt.Image = fileName
	}

	commentId, status, err := c.CreateCommentService(r.Context(), commnt)
	if err != nil {
		w.WriteHeader(status)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
		return
	}
	json.NewEncoder(w).Encode(struct {
		ID int `json:"id"`
	}{
		ID: commentId,
	})
}

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
