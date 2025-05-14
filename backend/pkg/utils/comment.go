package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"socialNetwork/entity"
)

// ParseAndValidateCommentForm parses and validates the comment form data.
func ParseAndValidateCommentForm(r *http.Request) (entity.Comment, []byte, error) {
	// Limit max memory to 10MB for multipart form
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return entity.Comment{}, nil, fmt.Errorf("error parsing form: %v", err)
	}

	var comment entity.Comment

	// Validate post_id (required)
	postIDStr := r.FormValue("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil || postID <= 0 {
		return comment, nil, fmt.Errorf("invalid or missing post_id")
	}
	comment.PostID = postID

	// Validate content (required)
	comment.Content = r.FormValue("content")
	if comment.Content == "" {
		return comment, nil, errors.New("content is required")
	}

	// Handle optional image
	var image []byte
	file, fileHeader, err := r.FormFile("image")
	if err == nil {
		image, comment.Image.NullString.String, err = ValidateImage(file, fileHeader)
		if err != nil {
			return comment, nil, fmt.Errorf("invalid image: %v", err)
		}
		comment.Image.SetValid(true)
	} else if err != http.ErrMissingFile {
		return comment, nil, fmt.Errorf("error reading uploaded file: %v", err)
	}

	return comment, image, nil
}
