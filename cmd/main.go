package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"snippetbox/cmd/web"
	"strconv"
)

// main calls run that initializes a http server
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Cannot load .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	web.Run(port)
}
