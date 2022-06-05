package contributors

import (
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
)


func CallBack(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/contributor/github/callback" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	tmpl := template.Must(template.ParseFiles("templates/app/contributors/callback.html"))
	tmpl.Execute(w, nil) 
	
}