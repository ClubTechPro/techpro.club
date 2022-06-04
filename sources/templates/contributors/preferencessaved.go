package contributors

import (
	"html/template"
	"net/http"
)


func PreferencesSaved(w http.ResponseWriter, r *http.Request){
	
	tmpl := template.Must(template.ParseFiles("templates/app/contributors/preferencessaved.html"))
	tmpl.Execute(w, nil) 
	
}
