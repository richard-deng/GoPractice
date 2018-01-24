package dlog

import (
	"os"
	"log"
	"fmt"
	"time"
)

var (
	Debug *log.Logger
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

func init() {
	logFile, err := os.OpenFile("./" + time.Now().Format("20060102") + ".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	Debug = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}