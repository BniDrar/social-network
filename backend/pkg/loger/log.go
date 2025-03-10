package loger

import (
	"io"
	"log"
	"os"
)

// Initloger initializes the logger
type CstmLogger struct {
	Info  *log.Logger
	Error *log.Logger
}

func NewLogger() *CstmLogger {
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		log.Panic("cannot create log file: ", err)
	}
	out := io.MultiWriter(file, os.Stdout)
	return &CstmLogger{
		Info:  log.New(out, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(out, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
