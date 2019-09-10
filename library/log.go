package library

import (
	log "github.com/sirupsen/logrus"
	. "os"
)

func InitializeLogging(logFile string) {
	log.SetLevel(log.InfoLevel)

	var file, err = OpenFile(logFile, O_RDWR|O_CREATE|O_APPEND, 0666)
	if err != nil {
		log.Warn("Could Not Open Log File : " + err.Error())
	}

	log.SetOutput(file)

	log.SetFormatter(&log.JSONFormatter{})
}
