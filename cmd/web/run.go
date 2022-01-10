package web

import (
	"net/http"
)

// Run starts a new http server with a default Multiplexer
func Run(address string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippets", ShowSnippets)
	mux.HandleFunc("/snippets/create", CreateSnippets)

	// Routes for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Initialize custom logger
	logger := NewLogger()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     address,
		ErrorLog: logger.errLog,
		Handler:  mux,
	}

	logger.infoLog.Printf("Starting server %s", address)
	err := srv.ListenAndServe()
	logger.errLog.Fatalln(err)
}
