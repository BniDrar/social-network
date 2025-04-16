package post

import (
	"database/sql"
	"fmt"
	"net/http"

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
	p.Service_GetAll(w, r)
}

func (p *post) ReactPost(w http.ResponseWriter, r *http.Request) {
	p.Service_React(w, r)
}

func (p *post) CreatePost(w http.ResponseWriter, r *http.Request) {
	// p.Service_CreateOne(w, r)
	err := FileUpload(r.FormFile("image"))
	fmt.Println(err)
}

func (p *post) GetPost(w http.ResponseWriter, r *http.Request) {
	p.Service_GetOne(w, r)
}

func (p *post) GetPostByUsername(w http.ResponseWriter, r *http.Request) {
	p.GetPostsByUserService(w, r)
}
