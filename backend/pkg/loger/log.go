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
		Info:  log.New(out, "INFO\t: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(out, "ERROR\t: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func NewTestLogger() *CstmLogger {
	return &CstmLogger{
		Info:  log.New(io.Discard, "", 0),
		Error: log.New(io.Discard, "", 0),
	}
}
