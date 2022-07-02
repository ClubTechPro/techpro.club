package projects

import (
	"net/http"

	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)

func Notifications(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/projects/notifications" {
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
}

func NotificationsCount(){

}