package web

import (
	"log"
	"os"
)

type logger struct {
	infoLog *log.Logger
	errLog  *log.Logger
}

func NewLogger() *logger {
	return &logger{
		log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime),
		log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
