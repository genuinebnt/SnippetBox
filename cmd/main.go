package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"snippetbox/cmd/web"
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
