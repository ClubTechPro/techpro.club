package projects

import (
	"html/template"
	"net/http"
)


func ProjectSaved(w http.ResponseWriter, r *http.Request){
	
	tmpl := template.Must(template.ParseFiles("templates/app/projects/projectsaved.html"))
	tmpl.Execute(w, nil) 
	
}
