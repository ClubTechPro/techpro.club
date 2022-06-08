package projects

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
)


func ProjectSaved(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/projects/thankyou" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	tmpl, err := template.New("").ParseFiles("templates/app/projects/projectsaved.html", "templates/app/projects/common/base.html")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "projectbase", nil) 
		}
	
}
