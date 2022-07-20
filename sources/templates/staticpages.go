package templates

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/users"
)

// Handles landing page
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
		} else {
			// Delete cookies
			users.DeleteSessionCookie(w, r)
			users.DeleteUserCookie(w, r)
		}
	}

	
	tmpl := template.Must(template.ParseFiles("templates/home/index.gohtml"))
    tmpl.Execute(w, nil)
}


// Handles Contact us page
func ContactUs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contactus" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	

	pageTitle := common.PageTitle{Title : "Contact Us"}

	tmpl, err := template.New("").ParseFiles("templates/home/contactus.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageTitle) 
	}
	
}

// Handles Careers page
func Careers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/careers" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	pageTitle := common.PageTitle{Title : "Careers"}

	tmpl, err := template.New("").ParseFiles("templates/home/careers.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageTitle) 
	}
}

// Handles Company page
func Company(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/company" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	pageTitle := common.PageTitle{Title : "About us"}

	tmpl, err := template.New("").ParseFiles("templates/home/company.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageTitle) 
	}
}

// Handles Brand page
func Brand(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/brand" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	pageTitle := common.PageTitle{Title : "The Brand"}

	tmpl, err := template.New("").ParseFiles("templates/home/brand.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageTitle) 
	}
}

// Handles Videos page
func Videos(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/videos" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	pageTitle := common.PageTitle{Title : "Training Videos"}

	tmpl, err := template.New("").ParseFiles("templates/home/videos.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageTitle) 
	}
}

// Page not found. 404 handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404. Page not found")
	}
}

