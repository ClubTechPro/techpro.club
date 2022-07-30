package pages

import (
	"net/http"

	"techpro.club/sources/users"
)

func Logout(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/logout" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Delete cookies
	users.DeleteSessionCookie(w, r)
	users.DeleteUserCookie(w, r)

	// Redirect to home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}