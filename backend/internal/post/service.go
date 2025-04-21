package post

import (
	"bufio"
	"context"
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
)

func (p *post) Service_GetAll(ctx context.Context) (data []byte, err error) {
	id := 1 // ctx.Value(entity.ContextID).(int)
	posts, err := p.Repo_GetAll(ctx, id)
	if err != nil {
		return
	}
	data, err = json.Marshal(posts)
	return
}

func (p *post) Service_GetOne(ctx context.Context, post_str string) (data []byte, err error) {
	id := 1 // ctx.Value(entity.ContextID).(int)
	post_id, err := strconv.Atoi(post_str)
	if err != nil {
		return
	}
	if p.Repo_UserCanPost(ctx, id, post_id) {
		var post entity.Post
		post, err = p.Repo_GetOne(ctx, post_id)
		if err != nil {
			return
		}
		data, err = json.Marshal(post)
		return
	} else {
		err = errors.New("you can't see post")
		return
	}
}

func (p *post) Service_CreateOne(ctx context.Context, body io.ReadCloser, users []string, post entity.Post) error {
	id := 1 // ctx.Value(entity.ContextID).(int)
	ImageFileName := post.Image
	json.NewDecoder(body).Decode(&post)
	post.Image = ImageFileName
	err := post.Validate(users)
	if err != nil {
		return errors.New(string("{ error: " + err.Error() + " }"))
	}
	p.Repo_CreatePost(ctx, id, post)
	return nil
}

func (p *post) Service_React(ctx context.Context, body io.ReadCloser) (err error) {
	id := 1 // ctx.Value(entity.ContextID).(int)
	react := entity.Vote{}
	json.NewDecoder(body).Decode(&react)
	if p.Repo_UserCanPost(ctx, id, react.ID) {
		err = p.Repo_React(ctx, id, react)
	} else {
		err = errors.New("you can't see the post")
	}
	return
}

func (p *post) GetPostsByUserService(ctx context.Context, username string) (data []byte, err error) {
	userID := 1 // ctx.Value(entity.ContextID).(int)
	posts, err := p.GetPostsByUserID(ctx, userID, username)
	if err != nil {
		return
	}
	data, err = json.Marshal(posts)
	if err != nil {
		return
	}
	return
}

func FileUpload(NeWFileName string, File multipart.File, FileHeader *multipart.FileHeader, err error) error {
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
	os.WriteFile("../"+NeWFileName+".png", FileContent, 0o444)
	return nil
}
