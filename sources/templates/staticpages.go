package templates

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/users"
)

// Handles landing page of Contributors
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	status, _ := users.GetSession(w,r)
	if status {
		sessionOk, _ := users.ValidateDbSession(w, r)
		if(sessionOk){
			http.Redirect(w, r, "/contributors/feeds", http.StatusSeeOther)
		}
	}

	
	tmpl := template.Must(template.ParseFiles("templates/home/index.html"))
    tmpl.Execute(w, nil)
}


// Handles landing page of Projects
func ProjectIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/projects" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	status, _ := users.GetSession(w,r)
	if status {
		sessionOk, _ := users.ValidateDbSession(w, r)
		if(sessionOk){
			http.Redirect(w, r, "/projects/list", http.StatusSeeOther)
		}
	}
		
		
	tmpl := template.Must(template.ParseFiles("templates/home/project_index.html"))
    tmpl.Execute(w, nil)
}

// Page not found. 404 handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404. Page not found")
	}
}

