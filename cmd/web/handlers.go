package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Home page route
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errLog.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.errLog.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
}

// CreateSnippets is Route that create snippets. Accepts only post requests
func (app *application) CreateSnippets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet"))
}

// ShowSnippets Shows created snippets based on id. id can be >= 1
func (app *application) ShowSnippets(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying snippets for id:%d", id)
}
