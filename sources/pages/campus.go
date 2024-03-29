package pages

import (
	"fmt"
	"html/template"
	"net/http"

	"techpro.club/sources/common"
)

// Function for Campus
func Campus(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/campus" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "Campus"}

	tmpl, err := template.New("").ParseFiles("templates/home/campus.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
}


// Function for CampusOnboard
func CampusOnboard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/campus/onboard" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title : "Campus Onboarding"}

	tmpl, err := template.New("").ParseFiles("templates/home/campusonboard.gohtml", "templates/home/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "basehome", pageDetails) 
	}
}