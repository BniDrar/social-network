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
	formData, image, err := utils.ParseAndValidateCommentForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	commentId, status, err := c.CreateCommentService(r.Context(), formData, image)
	if err != nil {
		c.loger.Error.Println(err)
		w.WriteHeader(status)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
		return
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(struct {
		ID int `json:"id"`
		Image string `json:"image"`
	}{
		ID: commentId,
		Image: formData.Image.NullString.String,
	})
}

func (c *comment) GetComments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || postId <= 0 {
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
