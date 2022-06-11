package contributors

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)

func Feeds(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/contributors/feeds" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	sessionOk, _ := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	
	tmpl, err := template.New("").ParseFiles("templates/app/contributors/feeds.gohtml", "templates/app/contributors/common/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", nil) 
	}

}