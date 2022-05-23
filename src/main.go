package main

import (
	"net/http"
	"text/template"
)

type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

func main() {
    fs := http.FileServer(http.Dir("../assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))


    tmpl := template.Must(template.ParseFiles("../templates/home/index.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, nil)
    })

    http.ListenAndServe(":8080", nil)
}