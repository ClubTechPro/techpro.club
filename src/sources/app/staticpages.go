package app

import (
	"html/template"
	"net/http"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/home/index.html"))
    tmpl.Execute(w, nil)
}
