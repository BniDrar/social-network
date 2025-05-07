package post

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/errors"
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
		errors.WriteError(w, errors.BadRequest("Method not allowed", nil))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	
	var cursor entity.Cursor
	if err := json.NewDecoder(r.Body).Decode(&cursor); err != nil {
		errors.WriteError(w, errors.BadRequest("Invalid request body", err))
		return
	}

	posts, status, err := p.GetPostsService(r.Context(), cursor)
	if err != nil {
		p.loger.Error.Println(err)
		errors.WriteError(w, err)
		return
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&posts)
}

func (p *post) GetGroupPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errors.WriteError(w, errors.BadRequest("Method not allowed", nil))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var cursor entity.Cursor
	if err := json.NewDecoder(r.Body).Decode(&cursor); err != nil {
		errors.WriteError(w, errors.BadRequest("Invalid request body", err))
		return
	}

	status, posts, err := p.getPostsByGroup(r.Context(), cursor)
	if err != nil {
		p.loger.Error.Println(err)
		errors.WriteError(w, err)
		return
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&posts)
}

func (p *post) ReactPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errors.WriteError(w, errors.BadRequest("Method not allowed", nil))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	postId, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil || postId == 0 {
		errors.WriteError(w, errors.BadRequest("Invalid post ID", err))
		return
	}

	status, err := p.PostEngagementService(r.Context(), postId)
	if err != nil {
		p.loger.Error.Println(err)
		errors.WriteError(w, err)
		return
	}

	w.WriteHeader(status)
}

func (p *post) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errors.WriteError(w, errors.BadRequest("Method not allowed", nil))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	formData, image, err := utils.ParseAndValidatePostForm(r)
	if err != nil {
		errors.WriteError(w, errors.BadRequest("Invalid form data", err))
		return
	}

	postId, status, err := p.CreatePostService(r.Context(), formData, image)
	if err != nil {
		p.loger.Error.Println(err)
		errors.WriteError(w, err)
		return
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(struct {
		PostID int `json:"post_id"`
	}{PostID: postId})
}

func (p *post) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errors.WriteError(w, errors.BadRequest("Method not allowed", nil))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	postId, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil {
		errors.WriteError(w, errors.BadRequest("Invalid post ID", err))
		return
	}

	post, status, err := p.GetPostService(r.Context(), postId)
	if err != nil {
		p.loger.Error.Println(err)
		errors.WriteError(w, err)
		return
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&post)
}
