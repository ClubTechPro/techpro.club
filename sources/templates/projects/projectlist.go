package projects

import (
	"html/template"
	"net/http"
)

func ProjectList(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("../templates/app/projects/projectlist.html"))
	tmpl.Execute(w, nil) 
}