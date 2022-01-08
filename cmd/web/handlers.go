package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// Home page route
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		files := []string{
			"./ui/html/home.page.tmpl",
			"./ui/html/base.layout.tmpl",
			"./ui/html/footer.partial.tmpl",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server error", 500)
			return
		}
		return
	}
	http.NotFound(w, r)
}

// CreateSnippets is Route that create snippets. Accepts only post requests
func CreateSnippets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet"))
}

// ShowSnippets Shows created snippets based on id. id can be >= 1
func ShowSnippets(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying snippets for id:%d", id)
}
