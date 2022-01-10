package web

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippets", app.ShowSnippets)
	mux.HandleFunc("/snippets/create", app.CreateSnippets)

	// Routes for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
