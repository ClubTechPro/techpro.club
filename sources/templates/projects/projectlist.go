package projects

import (
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
)

func ProjectList(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/projects/list" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	tmpl := template.Must(template.ParseFiles("templates/app/projects/projectlist.html"))
	tmpl.Execute(w, nil) 
}