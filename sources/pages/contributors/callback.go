package contributors

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/pages"
)


func CallBackGithub(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/contributors/github/callback" {
        pages.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	tmpl, err := template.ParseFiles("templates/app/contributors/callback.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.Execute(w, nil) 
	}
	
}

func CallBackGoogle(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/contributors/google/callback" {
        pages.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	tmpl, err := template.ParseFiles("templates/app/contributors/callback.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.Execute(w, nil) 
	}
	
	
}