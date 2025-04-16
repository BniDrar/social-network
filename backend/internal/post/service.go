package post

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"socialNetwork/entity"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/uuid"
)

func (p *post) Service_GetAll(w http.ResponseWriter, r *http.Request) {
	id := 1 // r.Context().Value(entity.ContextID).(int)
	posts, err := p.Repo_GetAll(r.Context(), id)
	if err != nil {
		w.Write([]byte(entity.WhereIsError() + " " + err.Error()))
		return
	}
	data, err := json.Marshal(posts)
	if err != nil {
		w.Write([]byte(entity.WhereIsError() + " " + err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (p *post) Service_GetOne(w http.ResponseWriter, r *http.Request) {
	id := 1 // r.Context().Value(entity.ContextID).(int)
	post_str := r.PathValue("id")
	post_id, err := strconv.Atoi(post_str)
	if err != nil {
		w.Write([]byte("id : " + post_str))
		w.Write([]byte(entity.WhereIsError() + " " + err.Error()))
		return
	}
	if p.Repo_UserCanPost(r.Context(), id, post_id) {
		post, err := p.Repo_GetOne(r.Context(), post_id)
		if err != nil {
			w.Write([]byte(entity.WhereIsError() + " " + err.Error()))
			return
		}
		data, err := json.Marshal(post)
		if err != nil {
			w.Write([]byte(entity.WhereIsError() + " " + err.Error()))
			return
		}
		w.Write(data)
	} else {
		w.Write([]byte("{ error: 'post not found' }"))
		return
	}
}

func (p *post) Service_CreateOne(w http.ResponseWriter, r *http.Request) {
	id := 1 // r.Context().Value(entity.ContextID).(int)
	post := entity.Post{}
	json.NewDecoder(r.Body).Decode(&post)
	err := post.Validate()
	if err != nil {
		w.Write([]byte("{ error: " + err.Error() + " }"))
	}
	p.Repo_CreatePost(r.Context(), id, post)
	w.Write([]byte(`{
		result: "done"
	}`))
}

func (p *post) Service_React(w http.ResponseWriter, r *http.Request) {
	id := 1 // r.Context().Value(entity.ContextID).(int)
	react := entity.Vote{}
	json.NewDecoder(r.Body).Decode(&react)
	if p.Repo_UserCanPost(r.Context(), id, react.ID) {
		err := p.Repo_React(r.Context(), id, react)
		if err != nil {
			w.Write([]byte(entity.WhereIsError() + " " + err.Error()))
			return
		}
		w.Write([]byte(`{
			result: "done"
		}`))
	} else {
		w.Write([]byte(" { error : 'you can't see the post' } "))
	}
}

// get
// 0 => table(group) contains user id
// 1 => table(follows) contains user id
// 2 => get direct

// react
// 0 => table(group) contains user id
// 1 => table(follows) contains user id
// if not status forbidden

func (p *post) GetPostsByUserService(w http.ResponseWriter, r *http.Request) {
	userID := 1 // r.Context().Value(entity.ContextID).(int)
	username := r.PathValue("username")
	posts, err := p.GetPostsByUserID(r.Context(), userID, username)
	if err != nil {
		w.Write([]byte(entity.WhereIsError() + " " + err.Error()))
		return
	}
	data, err := json.Marshal(posts)
	if err != nil {
		w.Write([]byte(entity.WhereIsError() + " " + err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func FileUpload(File multipart.File, FileHeader *multipart.FileHeader, err error) error {
	if err != nil {
		return err
	}
	fmt.Println((FileHeader.Size / 1024))
	var FileContent []byte
	var FileIsAccepted bool
	AcceptedTypes := []string{"png", "gif", "jpg", "jpeg"}
	for _, Type := range AcceptedTypes {
		if strings.HasSuffix(FileHeader.Filename, Type) {
			FileIsAccepted = true
		}
	}
	if !FileIsAccepted {
		return errors.New("file type is not accepted")
	}
	switch FileHeader.Filename[len(FileHeader.Filename)-3:] {
	case "gif":
		if (FileHeader.Size/1024) > 1024 || (FileHeader.Size/1024) < 10 {
			return errors.New("file size unmatched our conditions")
		}
	default:
		if (FileHeader.Size/1024) > 500 || (FileHeader.Size/1024) < 10 {
			return errors.New("file size unmatched our conditions")
		}
	}
	FileIsAccepted = false
	reader := bufio.NewReader(File)
	FileContent, _ = io.ReadAll(reader)
	AcceptedTypes = []string{"image/png", "image/jpg", "image/jpeg", "image/gif"}
	FileType := http.DetectContentType((FileContent))
	for _, Type := range AcceptedTypes {
		if Type == FileType {
			FileIsAccepted = true
		}
		fmt.Println(FileIsAccepted, " : ", Type, " == ", FileType)
	}
	if !FileIsAccepted {
		return errors.New("file type not matched")
	}
	NewFileName := uuid.NewString()
	os.WriteFile("../"+NewFileName+".png", FileContent, 0o444)
	return nil
}
