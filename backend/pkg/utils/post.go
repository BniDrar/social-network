package utils

import (
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
