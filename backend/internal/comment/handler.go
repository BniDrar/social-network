package comment

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/utils"
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
	w.Header().Set("Content-Type", "application/json")
	var commnt entity.Comment
	err := json.NewDecoder(r.Body).Decode(&commnt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err == nil {
		commnt.Image, err = utils.FileUpload(file, fileHeader)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
	}

	commentId, status, err := c.CreateCommentService(r.Context(), commnt)
	if err != nil {
		c.loger.Error.Println(err)
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
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || postId <= 0{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid post id"})
	}
	w.Header().Set("Content-Type", "application/json")
	status, commnts, err := c.GetCommentsService(r.Context(), postId)
	if err != nil {
		c.loger.Error.Println(err)
		w.WriteHeader(status)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
		return
	}
	json.NewEncoder(w).Encode(commnts)
	w.WriteHeader(status)
}
