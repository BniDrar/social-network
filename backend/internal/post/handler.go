package post

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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
	ServeMedia(w http.ResponseWriter, r *http.Request)
}

type post struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

func Newpost(dep *config.Dependencies) Post {
	return &post{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

func (p *post) ServeMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving media file:", r.URL.Path)

	// Clean the path to avoid path traversal
	path := filepath.Clean("." + r.URL.Path)[len("api/pictures/"):]
	fmt.Println(path)
	// Check if file exists and is not a directory
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, path)
}

func (p *post) GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "bad query"})
		return
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "bad query"})
		return
	}
	posts, status, err := p.GetPostsService(r.Context(), limit, offset)
	w.WriteHeader(status)
	if err != nil {
		p.loger.Error.Println(err)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
		return
	}
	json.NewEncoder(w).Encode(&posts)
}

func (p *post) ReactPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	postId, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil || postId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid query"})
	}
	status, err := p.PostEngagementService(r.Context(), postId)
	w.WriteHeader(status)
	if err != nil {
		p.loger.Error.Println(err)
	}
}

func (p *post) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	formData, image, err := utils.ParseAndValidatePostForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}

	postId, status, err := p.CreatePostService(r.Context(), formData, image)
	w.WriteHeader(status)
	if err != nil {
		p.loger.Error.Println(err)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid credentials"})
		}
		return
	}

	json.NewEncoder(w).Encode(struct {
		PostID int `json:"post_id"`
	}{PostID: postId})
}

func (p *post) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	postId, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "invalid query"})
		return
	}
	post, status, err := p.GetPostService(r.Context(), postId)
	w.WriteHeader(status)
	if err != nil {
		p.loger.Error.Println(err)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
		return
	}
	json.NewEncoder(w).Encode(&post)
}
