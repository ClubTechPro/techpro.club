package contributors

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/pages/videos"
	"techpro.club/sources/users"
)

func VideoFeeds(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contributors/videofeeds" {
        pages.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	// Fetch notificaitons
	_, _, notificationsCount, notificationsList := pages.NotificationsCountAndTopFive(userID)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}
		

	if r.Method == "GET"{

		pageTitle := common.PageTitle{Title : "Video Feeds"}

		output := videos.FinalVideoListOutStruct{
			userNameImage,
			notificationsCount,
			notificationsList,
			pageTitle,
		}

		tmpl, err := template.New("").ParseFiles("templates/app/common/base.gohtml", "templates/app/common/contributormenu.gohtml",  "templates/app/contributors/videofeeds.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "base", output) 
		}

	}
}