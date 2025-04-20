package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
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
