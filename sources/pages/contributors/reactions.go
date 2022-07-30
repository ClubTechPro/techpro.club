package contributors

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/users"
)

// Fetched reacted projects

func FetchReactions(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contributors/reactions" {
		pages.ErrorHandler(w, r, http.StatusNotFound)
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

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	fmt.Println(userNameImage)


	tmpl, err := template.ParseFiles("templates/app/contributors/common/base.gohtml", "templates/app/contributors/reactions.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", nil) 
	}
}

// Fetch all reacted projects
// func fetchReactedProjects(userID primitive.ObjectID){}
