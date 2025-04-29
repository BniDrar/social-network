package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"unicode"

	"socialNetwork/entity"
)

func CheckCommentsCredentials(comnt entity.Comment) bool {
	return len(removeAllWhitespace(comnt.Content)) > 0 &&
		len(removeAllWhitespace(comnt.Content)) < 200 &&
		comnt.PostID > 0
}

func removeAllWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func ValidateImage(file multipart.File, fileHeader *multipart.FileHeader) ([]byte, string, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, "", err
	}

	mimeType := http.DetectContentType(content)
	allowedMimes := map[string]string{
		"image/png":  "png",
		"image/jpeg": "jpg",
		"image/jpg":  "jpg", // technically not needed, "image/jpeg" covers jpg/jpeg
		"image/gif":  "gif",
	}
	ext, ok := allowedMimes[mimeType]
	if !ok {
		return nil, "", errors.New("unsupported MIME type")
	}

	sizeKB := len(content) / 1024
	if sizeKB > 1024 || sizeKB < 10 {
		return nil, "", errors.New("file size invalid")
	}

	// Generate secure random filename
	name, err := randomHex(12)
	if err != nil {
		return nil, "", err
	}
	path := filepath.Join("./media", name+"."+ext)

	return content, path, nil
}

// Helper: Generate random hex string
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
