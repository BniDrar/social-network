package comment

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type comment struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

type Comment interface {
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Vote(w http.ResponseWriter, r *http.Request)
}

/*               app     */
func NewComment(dep *config.Dependencies) Comment {
	return &comment{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

func (c *comment) Post(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	prep, err := c.db.PrepareContext(r.Context(), ``)
	if err != nil {
		return
	}
	comment := entity.Comment{}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		return
	}
	_, err = prep.ExecContext(r.Context(), id, comment.Content, comment.PostID, time.Now())
	if err != nil {
		return
	}
	w.Write([]byte(`{
		result:"done"
	}`))
}

func (c *comment) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	prep, err := c.db.PrepareContext(r.Context(), ``)
	if err != nil {
		return
	}
	comment := entity.Comment{}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		return
	}
	comments := []entity.Comment{}
	res, err := prep.QueryContext(r.Context(), id, comment.PostID)
	if err != nil {
		return
	}
	for res.Next() {
		comment := entity.Comment{}
		err = res.Scan(&comment.ID, &comment.UserID, &comment.Content, &comment.PostID)
		if err != nil {
			return
		}
		comments = append(comments, comment)
	}
	data, err := json.Marshal(comments)
	if err != nil {
		return
	}
	w.Write(data)
}

func (c *comment) Vote(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	prep, err := c.db.PrepareContext(r.Context(), ``)
	if err != nil {
		return
	}
	react := entity.CommentReaction{}
	err = json.NewDecoder(r.Body).Decode(&react)
	if err != nil {
		return
	}
	_, err = prep.Exec(id, react.ID, react.Status)
	if err != nil {
		return
	}
	w.Write([]byte(`{
		result: "done"
	}`))
}
