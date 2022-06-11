package contributors

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
)


func CallBack(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/contributors/github/callback" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	tmpl, err := template.New("").ParseFiles("templates/app/contributors/callback.gohtml", "templates/app/contributors/common/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", nil) 
	}
	
	
}