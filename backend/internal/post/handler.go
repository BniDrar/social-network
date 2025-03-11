package post

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type Post interface {
	Post(w http.ResponseWriter, r *http.Request)
	// GetPosts(w http.ResponseWriter, r *http.Request)
	// React(w http.ResponseWriter, r *http.Request)
}

type post struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

var _ string = `{
					"id": 1,
					"title": "My First Post",
					"image": "iVBORw0KGgoAAAANSUhEUgAA...kJggg==",
					"content": "This is the content of the post.",
					"user_id": 42,
					"status": 1,
					"group": 5,
					"created_at": "2025-03-11T12:34:56Z",
					"updated_at": "2025-03-11T12:34:56Z"
				}`

func Newpost(dep *config.Dependencies) Post {
	return &post{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

func (u *post) GetPosts(w http.ResponseWriter, r *http.Request) {
	// id := r.Context().Value(entity.ContextID).(int)
	// prep, err := u.db.PrepareContext(r.Context(), ``)
	// if err != nil {
	// 	return
	// }
	// posts := []entity.Post{}
	// res, err := prep.QueryContext(r.Context(), id)
	// if err != nil {
	// 	return
	// }
	// for res.Next() {
	// 	post := entity.Post{}
	// 	err := res.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.GroupID)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	posts = append(posts, post)
	// }
	// data, err := json.Marshal(posts)
	// if err != nil {
	// 	return
	// }
	// w.Write(data)
}

/*             NextJs                */
func (u *post) React(w http.ResponseWriter, r *http.Request) {
	// id := r.Context().Value(entity.ContextID).(int)
	// prep, err := u.db.PrepareContext(r.Context(), ``)
	// if err != nil {
	// 	return
	// }
	// react := entity.PostReaction{}
	// err = json.NewDecoder(r.Body).Decode(&react)
	// if err != nil {
	// 	return
	// }
	// _, err = prep.Exec(id, react.ID, react.Status)
	// if err != nil {
	// 	return
	// }
	// w.Write([]byte(`{
	// 	result: "done"
	// }`))
}

func (u *post) Post(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case "POST":
	// 	u.CreatePost(w, r)
	// case "GET":
	// 	u.GetPost(w, r)
	// default:
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// }
}

// func (u *post) CreatePost(w http.ResponseWriter, r *http.Request) {
// 	id := r.Context().Value(entity.ContextID).(int)
// 	prep, err := u.db.PrepareContext(r.Context(), ``)
// 	if err != nil {
// 		return
// 	}
// 	post := entity.Post{}
// 	json.NewDecoder(r.Body).Decode(&post)
// 	_, err = prep.Exec(id, post.Title, post.Content, post.Image, post.GroupID, post.Status)
// 	if err != nil {
// 		return
// 	}
// 	w.Write([]byte(`{
// 		result: "done"
// 	}`))
// }

func (p *post) Create(w http.ResponseWriter, r *http.Request) {
	user_id, exist := r.Context().Value(entity.ContextID).(int)
	if !exist {
		p.loger.Error.Println("user doesn't exists")
		w.WriteHeader(http.StatusForbidden) 
		return
	}

	post := entity.Post{}
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post_id, status, err := p.CreatePostService(post, user_id)
	if err != nil {
		w. WriteHeader(status)
		return
	}

	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("{id: %v}", post_id)))
}

func (p *post) GetPostById (w http.ResponseWriter, r *http.Request) {
	// UserID, exist:= r.Context().Value(entity.ContextID).(int)
	// if !exist {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }

}

func (u *post) GetPost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	prep, err := u.db.PrepareContext(r.Context(), ``)
	if err != nil {
		return
	}
	post := entity.Post{}
	res := prep.QueryRowContext(r.Context(), id)
	err = res.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Status, &post.GroupID)
	if err != nil {
		return
	}
	data, err := json.Marshal(post)
	if err != nil {
		return
	}
	w.Write(data)
}
