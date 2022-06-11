package projects

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)

func ProjectPreview(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/projects/view" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/projects", http.StatusSeeOther)
	}

	projectID := r.URL.Query().Get("projectid")
	result := FetchProjectDetails(projectID, userID)
	

	tmpl, err := template.New("").ParseFiles("templates/app/projects/projectpreview.gohtml", "templates/app/projects/common/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "projectbase", result) 
	}
}