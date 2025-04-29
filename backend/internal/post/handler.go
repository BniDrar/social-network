package post

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/utils"
	"socialNetwork/pkg/websocket"
)

type Post interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
	ReactPost(w http.ResponseWriter, r *http.Request)
	GetPostByUsername(w http.ResponseWriter, r *http.Request)
}

type post struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

func Newpost(dep *config.Dependencies) Post {
	return &post{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

func (p *post) GetPosts(w http.ResponseWriter, r *http.Request) {
	data, err := p.Service_GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	w.Write(data)
}

func (p *post) ReactPost(w http.ResponseWriter, r *http.Request) {
	err := p.Service_React(r.Context(), r.Body)
	if err != nil {
	}
}

func (p *post) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var (
		post  entity.Post
		image []byte
	)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid response body"})
	}
	file, fileHeader, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Error reading uploaded file: " + err.Error()})
		return
	}
	if err != nil {
		post.Image = ""
	} else {
		image, post.Image, err = utils.ValidateImage(file, fileHeader)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
			return
		}
	}
	status, postId, err := p.CreatePostService(r.Context(), post, image)
	if err != nil {
		w.WriteHeader(status)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid Credentials"})
		}
		return
	}
	json.NewEncoder(w).Encode(struct {
		PostID int `json:"post_id"`
	}{
		PostID: postId,
	})
	w.WriteHeader(status)	
}

func (p *post) GetPost(w http.ResponseWriter, r *http.Request) {
	data, err := p.Service_GetOne(r.Context(), r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.Write(data)
}

func (p *post) GetPostByUsername(w http.ResponseWriter, r *http.Request) {
	data, err := p.GetPostsByUserService(r.Context(), r.PathValue("username"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.Write(data)
}
