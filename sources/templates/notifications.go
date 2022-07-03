package templates

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

	var Notificaitons []common.FetchNotificationStruct
	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, userName, image := FetchUsernameImage(w, r)

	if(!status){
		log.Println("Error fetching user name and image from cookies")
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchNotifications := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_NOTIFICATIONS)
	notifications, errNotifications := fetchNotifications.Find(context.TODO(), bson.M{"userid" : userID})

	if errNotifications != nil{
		fmt.Println(errNotifications.Error())
	} else {
		for notifications.Next(context.TODO()){
			var notification common.FetchNotificationStruct
			errNotifications := notifications.Decode(&notification)

			if errNotifications != nil{
				fmt.Println(errNotifications.Error())
			} else {
				Notificaitons = append(Notificaitons, notification)
			}
		}
	}

	output := NotificationStruct{Notificaitons, userNameImage}

	tmpl, err := template.New("").ParseFiles("templates/app/notifications.gohtml", "templates/app/contributors/common/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", output) 
	}
}

// Find total unread notifications for a user from database
func NotificationsCount(userID primitive.ObjectID)(status bool, msg string, count int64){
	status = true
	msg = "Success"
	count = 0

	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	countNotifications := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_NOTIFICATIONS)
	count, errCount := countNotifications.CountDocuments(context.TODO(), bson.M{"read": false, "userid" : userID})

	if errCount != nil{
		status = false
		msg = errCount.Error()
	} else {
		status = true
		msg = ""
	}

	return status, msg, count
}