package entity

import (
	"errors"
	"image"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func isValidUploadedImage(file multipart.File, header *multipart.FileHeader) (bool, error) {
	defer file.Seek(0, io.SeekStart)

	if header.Size > 500*1024 {
		return false, errors.New("File is too large (must be < 500KB).")
	}

	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return false, errors.New("Unable to read file.")
	}

	contentType := http.DetectContentType(buff)
	if !strings.HasPrefix(contentType, "image/") {
		return false, errors.New("File is not an image.")
	}
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif" {
		return false, errors.New("Only PNG, JPG, and GIF files are allowed.")
	}

	_, _, err := image.DecodeConfig(file)
	if err != nil {
		return false, errors.New("Invalid image file.")
	}
	return true, errors.New("Image is valid.")
}