package entity

import (
	"runtime"
	"strconv"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WhereIsError() string {
	_, file, line, _ := runtime.Caller(1)
	return file + ":" + strconv.Itoa(line)
}
