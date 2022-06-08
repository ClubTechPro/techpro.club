package contributors

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
)

func Feeds(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/contributor/feeds" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	tmpl, err := template.New("").ParseFiles("templates/app/contributors/feeds.html", "templates/app/contributors/common/base.html")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", nil) 
	}

}