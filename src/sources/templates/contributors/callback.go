package contributors

import (
	"html/template"
	"net/http"
)


func CallBack(w http.ResponseWriter, r *http.Request){
	
		tmpl := template.Must(template.ParseFiles("../templates/app/contributors/callback.html"))
		tmpl.Execute(w, nil) 
	
}