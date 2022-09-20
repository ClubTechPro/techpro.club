package pages

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

	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "Contact Us"}

	tmpl, err := template.New("").ParseFiles("templates/home/contactus.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
	
}

// Handles Careers page
func Careers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/careers" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "Careers"}

	tmpl, err := template.New("").ParseFiles("templates/home/careers.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
}

// Handles Company page
func Company(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/company" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "About us"}

	tmpl, err := template.New("").ParseFiles("templates/home/company.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
}

// Handles Brand page
func Brand(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/brand" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "The Brand"}

	tmpl, err := template.New("").ParseFiles("templates/home/brand.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
}

// Handles Videos page
func Videos(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/videos" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "Training Videos"}

	tmpl, err := template.New("").ParseFiles("templates/home/videos.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
}

// Privacy policy page
func PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/privacy-policy" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "Privacy Policy"}

	tmpl, err := template.New("").ParseFiles("templates/home/privacy.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
}

// Cookie policy page
func CookiePolicy(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cookie-policy" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "Cookie Policy"}

	tmpl, err := template.New("").ParseFiles("templates/home/cookie.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
}

// Cookie policy page
func TermsOfService(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/terms-and-conditions" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }	
		
	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "Terms and Conditions"}

	tmpl, err := template.New("").ParseFiles("templates/home/terms.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
}

// Page not found. 404 handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "Page not found"}

		tmpl, err := template.New("").ParseFiles("templates/home/404.gohtml", "templates/home/base.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
		}
	}
}

