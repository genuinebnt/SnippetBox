package web

import (
	"fmt"
	"log"
	"net/http"
)

// Run starts a new http server with a default Multiplexer
func Run(port int) {
	address := fmt.Sprintf("localhost:%d", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippets", ShowSnippets)
	mux.HandleFunc("/snippets/create", CreateSnippets)

	// Routes for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Listening to port", port)
	err := http.ListenAndServe(address, mux)
	log.Fatalln(err)
}
