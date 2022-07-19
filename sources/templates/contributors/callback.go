package contributors

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/templates"
)


func CallBack(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/contributors/github/callback" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := templates.FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	tmpl, err := template.New("").ParseFiles("templates/app/contributors/callback.gohtml", "templates/app/contributors/common/base_new.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", userNameImage) 
	}
	
	
}