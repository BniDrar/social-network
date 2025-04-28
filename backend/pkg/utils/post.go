package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"unicode"

	"socialNetwork/entity"

	"github.com/google/uuid"
)

var CompletePath = "./media/"

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

func FileUpload(File multipart.File, FileHeader *multipart.FileHeader, backErr error) (FilePath string, err error) {
	if backErr != nil {
		err = backErr
		return
	}
	fmt.Println(FileHeader.Filename)
	fmt.Println(FileHeader.Size)
	fmt.Println(FileHeader.Header)
	fmt.Println(FileHeader.Size/1024, "KB")
	var FileContent []byte
	var FileIsAccepted bool
	AcceptedTypes := []string{"png", "gif", "jpg", "jpeg"}
	for _, Type := range AcceptedTypes {
		if strings.HasSuffix(FileHeader.Filename, Type) {
			FileIsAccepted = true
		}
	}
	if !FileIsAccepted {
		err = errors.New("file type is not accepted")
		return
	}
	if (FileHeader.Size/1024) > 1024 || (FileHeader.Size/1024) < 10 {
		err = errors.New("file size unmatched our conditions")
		return
	}
	FileIsAccepted = false
	reader := bufio.NewReader(File)
	FileContent, _ = io.ReadAll(reader)
	AcceptedTypes = []string{"image/png", "image/jpg", "image/jpeg", "image/gif"}
	FileType := http.DetectContentType((FileContent))
	fmt.Println("FileType : ", FileType)
	for _, Type := range AcceptedTypes {
		if Type == FileType {
			FileIsAccepted = true
		}
		fmt.Println(FileIsAccepted, " : ", Type, " == ", FileType)
	}
	if !FileIsAccepted {
		err = errors.New("file type not matched")
	}
	if err == nil {
		FilePath = CompletePath + uuid.NewString()
	}
	return
}
