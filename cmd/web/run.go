package web

import (
	"net/http"
)

// Run starts a new http server with a default Multiplexer
func Run(address string) {
	// Initialize custom logger
	app := NewApp()

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippets", app.ShowSnippets)
	mux.HandleFunc("/snippets/create", app.CreateSnippets)

	// Routes for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     address,
		ErrorLog: app.errLog,
		Handler:  mux,
	}

	app.infoLog.Printf("Starting server %s", address)
	err := srv.ListenAndServe()
	app.errLog.Fatalln(err)
}
