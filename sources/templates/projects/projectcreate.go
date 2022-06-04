package projects

import (
	"html/template"
	"net/http"
)

func ProjectCreate(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("../templates/app/projects/projectcreate.html"))
	tmpl.Execute(w, nil) 
}