package pages

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
	"techpro.club/sources/users"
)


type NotificationStruct struct{
	Notifications []common.FetchNotificationStruct `json:"notifications"`
	UserNameImage common.UsernameImageStruct `json:"usernameImage"`
}

func Notifications(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/users/notifications" {
        ErrorHandler(w, r, http.StatusNotFound)
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
	status, msg, userName, image := FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	_, _, Notifications := fetchNotificationList(w, r, userID)

	output := NotificationStruct{Notifications, userNameImage}

	tmpl, err := template.New("").ParseFiles("templates/app/notifications.gohtml", "templates/app/contributors/common/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", output) 
	}
}

// Fetch notifications list
func fetchNotificationList(w http.ResponseWriter, r *http.Request, userID primitive.ObjectID)(status bool, msg string, Notifications []common.FetchNotificationStruct){
	status = false
	msg = ""

	status, msg, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchNotifications := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_NOTIFICATIONS)
	notifications, errNotifications := fetchNotifications.Find(context.TODO(), bson.M{"userid" : userID})

	if errNotifications != nil{
		fmt.Println(errNotifications.Error())
		msg = errNotifications.Error()
	} else {
		for notifications.Next(context.TODO()){
			var notification common.FetchNotificationStruct
			errNotifications := notifications.Decode(&notification)

			if errNotifications != nil{
				status = false
				msg = errNotifications.Error()
			} else {
				status = true
				msg = "Success"
				Notifications = append(Notifications, notification)
			}
		}
	}

	return status, msg, Notifications
}