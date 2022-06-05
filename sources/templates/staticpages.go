package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

// Handles landing page of Contributors
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	tmpl := template.Must(template.ParseFiles("templates/home/index.html"))
    tmpl.Execute(w, nil)
}


// Handles landing page of Projects
func ProjectIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/project" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	tmpl := template.Must(template.ParseFiles("templates/home/index.html"))
    tmpl.Execute(w, nil)
}

// Page not found. 404 handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404. Page not found")
	}
}

