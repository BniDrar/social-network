package post

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
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

var _ string = `{
  				  "id": "1",
  				  "avatar": "/assets/images/no-face.jpg",
  				  "username": "yrahhaou",
  				  "content": "🍲🍖 Classic Homestyle Meatloaf with a Tangy Glaze 🍋✨",
  				  "image": "/assets/images/post1.jpg",
  				  "likes": "15k",
  				  "comments": "5k",
  				  "user_like": false,
  				  "created_at": "7h"
  				}`

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
	File, FileHeader, err := r.FormFile("image")
	FilePath, err := FileUpload(File, FileHeader, err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "File Not Uploaded"})
		return
	}
	post := entity.Post{
		Image: FilePath,
	}
	err = p.Service_CreateOne(r.Context(), r.Body, r.Form["users"], post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	Content, _ := io.ReadAll(File)
	os.WriteFile(FilePath, Content, 0o444)
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
