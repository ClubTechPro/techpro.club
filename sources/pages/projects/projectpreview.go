package projects

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/users"
)

type FinalProjectPreviewOutStruct struct{
	ProjectPreview common.FetchProjectStruct `json:"projectsList"`
	UserNameImage common.UsernameImageStruct `json:"userNameImage"`
	ProjectOwner bool `json:"projectOwner"`
	MyBookmarks []primitive.ObjectID `json:"myBookmarks"`
	MyReactions []primitive.ObjectID `json:"myReactions"`
	NotificaitonsCount int64 `json:"notificationsCount"`
	NotificationsList []common.MainNotificationStruct `json:"nofiticationsList"`
	PageTitle common.PageTitle `json:"pageTitle"`
}

func ProjectPreview(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/projects/view" {
        pages.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/projects", http.StatusSeeOther)
	}

	var finalOutStruct FinalProjectPreviewOutStruct
	var userNameImage common.UsernameImageStruct

	// Fetch notificaitons
	_, _, notificationsCount, notificationsList := pages.NotificationsCountAndTopFive(userID)

	// Fetch reactions and bookmarks
	_, _, bookmarks, reactions := pages.FetchMyBookmarksAndReactions(userID)

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	var functions = template.FuncMap{
		"objectIdToString" : pages.ObjectIDToString,
		"containsObjectId" : pages.ContainsObjectID,
		"timeElapsed" : pages.TimeElapsed,
	}

	projectID := r.URL.Query().Get("projectid")
	projectStatus, projectError, result := pages.FetchProjectDetails(projectID, userID)

	projectOwner := false

	pageTitle := common.PageTitle{Title : result.ProjectName}
	

	if(projectStatus){

		if result.UserID == userID {
			projectOwner = true
		}

		finalOutStruct = FinalProjectPreviewOutStruct{result, userNameImage, projectOwner, bookmarks, reactions, notificationsCount, notificationsList, pageTitle}
	} else {
		fmt.Println(projectError)
		pages.ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	
	

	tmpl, err := template.New("").Funcs(functions).ParseFiles("templates/app/common/base.gohtml", "templates/app/common/projectmenu.gohtml", "templates/app/projects/projectpreview.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "base", finalOutStruct) 
	}
}