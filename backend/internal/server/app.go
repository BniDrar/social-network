package server

import (
	"io"
	"log"
	"os"
	config "socialNetwork/pkg/config"
)

const secret string = "Forum01Oujda"

func Run(cfg *config.Conf)  {
	// Prepare logger
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		log.Printf("cannot create log file: %v", err)
	}
	defer file.Close()
	logWriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logWriter)

}
