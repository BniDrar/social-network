package utils

import (
	"fmt"
	"net/http"
	"socialNetwork/entity"
	"strconv"
)

func ParseAndValidatePostForm(r *http.Request) (entity.Post, []byte, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("mother fucker")
		return entity.Post{}, nil, fmt.Errorf("error parsing form")
	}

	var post entity.Post
	post.Content = r.FormValue("content")
	if post.Content == "" {
		return post, nil, fmt.Errorf("content is required")
	}

	status, err := strconv.Atoi(r.FormValue("status"))
	if err != nil ||status <0 || status > 2 {
		return post, nil, fmt.Errorf("invalid status value")
	}
	post.Status = status

	if status == entity.PostStatusCustom {
		viewers, err := parseAllowedViewers(r.Form["allowed_viewers"])
		if err != nil {
			return post, nil, err
		}
		post.AllowedViewers = viewers
	}

	var image []byte
	file, fileHeader, err := r.FormFile("image")
	if err == nil {
		image, post.Image.NullString.String, err = ValidateImage(file, fileHeader)
		if err != nil {
			return post, nil, fmt.Errorf("invalid image: %v", err)
		}
		post.Image.SetValid(true)
	} else if err != http.ErrMissingFile {
		return post, nil, fmt.Errorf("error reading uploaded file: %v", err)
	}

	return post, image, nil
}

func parseAllowedViewers(values []string) ([]int, error) {
	var viewers []int
	for _, v := range values {
		id, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid allowed_viewer id: %v", v)
		}
		viewers = append(viewers, id)
	}
	if len(viewers) == 0 {
		return nil, fmt.Errorf("allowed_viewers required when status is custom")
	}
	return viewers, nil
}
