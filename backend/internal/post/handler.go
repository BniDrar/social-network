package post

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	GetGroupPosts(w http.ResponseWriter, r *http.Request)
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

	if len(r.URL.Path) >= len("/api/pictures/") || !strings.HasPrefix(r.URL.Path, "/api/pictures/") {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Clean the path to avoid path traversal
	path := filepath.Clean("." + r.URL.Path)[len("api/pictures/"):]
	// Check if file exists and is not a directory
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	http.ServeFile(w, r, path)
}

func (p *post) GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var cursor entity.Cursor
	err := json.NewDecoder(r.Body).Decode(&cursor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	posts, status, err := p.GetPostsService(r.Context(), cursor)
	w.WriteHeader(status)
	if err != nil {
		p.loger.Error.Println(err)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
		return
	}

	err = json.NewEncoder(w).Encode(&posts)
	if err != nil {
		p.loger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p *post) GetGroupPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var cursor entity.Cursor
	err := json.NewDecoder(r.Body).Decode(&cursor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}

	status, posts, err := p.getPostsByGroup(r.Context(), cursor)
	w.WriteHeader(status)
	if err != nil {
		p.loger.Error.Println(err)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(&posts)
}

func (p *post) ReactPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

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
