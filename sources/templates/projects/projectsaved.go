package projects

import (
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
)


func ProjectSaved(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/projects/thankyou" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	tmpl := template.Must(template.ParseFiles("templates/app/projects/projectsaved.html"))
	tmpl.Execute(w, nil) 
	
}
