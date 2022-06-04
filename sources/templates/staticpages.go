package templates

import (
	"html/template"
	"net/http"
)

// Handles landing page of Contributors
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("../templates/home/index.html"))
    tmpl.Execute(w, nil)
}


// Handles landing page of Projects
func ProjectIndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/home/index.html"))
    tmpl.Execute(w, nil)
}
