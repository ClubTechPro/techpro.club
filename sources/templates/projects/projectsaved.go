package projects

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)


func ProjectSaved(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/projects/thankyou" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	sessionOk, _ := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/projects", http.StatusSeeOther)
	}
	
	tmpl, err := template.New("").ParseFiles("templates/app/projects/projectsaved.html", "templates/app/projects/common/base.html")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "projectbase", nil) 
		}
	
}
