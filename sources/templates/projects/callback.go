package projects

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
)


func CallBack(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/projects/github/callback" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	tmpl, err := template.New("").ParseFiles("templates/app/projects/callback.html", "templates/app/projects/common/base.html")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "projectbase", nil) 
	}
	
	
}