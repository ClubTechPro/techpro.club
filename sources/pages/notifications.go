package pages

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"techpro.club/sources/common"
	"techpro.club/sources/users"
)


type NotificationStruct struct{
	Notifications []common.FetchNotificationStruct `json:"notifications"`
	UserNameImage common.UsernameImageStruct `json:"usernameImage"`
	NotificaitonsCount int64 `json:"notificationsCount"`
	NotificationsList []common.MainNotificationStruct `json:"nofiticationsList"`
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

	// Fetch notificaitons
	_, _, notificationsCount, notificationsList := NotificationsCountAndTopFive(userID)

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	_, _, Notifications := fetchNotificationList(w, r, userID)

	output := NotificationStruct{Notifications, userNameImage, notificationsCount, notificationsList}

	tmpl, err := template.New("").ParseFiles("templates/app/notifications.gohtml", "templates/app/contributors/common/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", output) 
	}
}

// Mark notification read
func MarkNotificationRead(w http.ResponseWriter, r *http.Request){

	status := false
	msg := ""

	if r.URL.Path != "/api/marknotificationsread" {
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

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchNotifications := client.Database(dbName).Collection(common.CONST_MO_NOTIFICATIONS)
	errNotifications := fetchNotifications.FindOneAndUpdate(context.TODO(), bson.M{"userid" : userID}, 
													bson.M{"$set" : bson.M{"notificationsList.$[elem].read" : true, "unreadnotifications" : 0}}, 
													options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
														Filters: []interface{}{bson.M{"elem.read": false}},
													}),
												)

	if errNotifications.Err() != nil{
		msg = errNotifications.Err().Error()
	} else {
		status = true
		msg = "Success"
	}


	output := common.JsonOutput{
		Status: status,
		Msg: msg,
	}

	out, _ := json.Marshal(output)
	w.Write(out)
}

// Fetch notifications list
func fetchNotificationList(w http.ResponseWriter, r *http.Request, userID primitive.ObjectID)(status bool, msg string, Notifications []common.FetchNotificationStruct){
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchNotifications := client.Database(dbName).Collection(common.CONST_MO_NOTIFICATIONS)
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