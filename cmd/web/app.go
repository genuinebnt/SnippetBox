package web

import (
	"log"
	"os"
)

type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
}

func NewApp() *application {
	return &application{
		log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime),
		log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
