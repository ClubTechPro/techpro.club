package contributors

import (
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
)


func PreferencesSaved(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/contributors/thankyou" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	tmpl := template.Must(template.ParseFiles("templates/app/contributors/preferencessaved.html"))
	tmpl.Execute(w, nil) 
	
}
