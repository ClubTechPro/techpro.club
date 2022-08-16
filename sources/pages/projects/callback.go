package projects

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/pages"
)

func CallBack(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/projects/github/callback" {
		pages.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	if !status {
		log.Println(msg)
	} else {
		userNameImage = common.UsernameImageStruct{Username: userName, Image: image}
	}

	tmpl, err := template.New("").ParseFiles("templates/app/projects/callback.gohtml", "templates/app/projects/common/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		tmpl.ExecuteTemplate(w, "projectbase", userNameImage)
	}

}
