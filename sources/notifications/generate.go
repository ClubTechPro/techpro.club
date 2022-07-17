package notifications

import (
	"context"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
)

type Newtest struct {
	Msg    string
	Status bool
}

// Create a notification
func Create(w http.ResponseWriter, r *http.Request) {

	oid := primitive.NewObjectID()

	_, _, result := insert(oid, "contributor", "test subject", "test mesg", "https://link", false, false)

	tmpl, err := template.New("").ParseFiles("templates/app/notifications.gohtml", "templates/app/projects/common/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		tmpl.ExecuteTemplate(w, "projectbase", result)
	}

}


//fetch a notification
func Fetch(w http.ResponseWriter, r *http.Request) {

	_, _, result := FetchNotification()

	tmpl, err := template.New("").ParseFiles("templates/app/notifications.gohtml", "templates/app/projects/common/base.gohtml")

	fmt.Println("trying to execute")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
		tmpl.ExecuteTemplate(w, "projectbase", result)
	}
	

}

func insert(userId primitive.ObjectID, notificationType, subject, message, link string, read, isProject bool) (status bool, msg string, result Newtest) {
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	notificationCollection := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_NOTIFICATIONS)

	if isProject {
		notificationCollection = client.Database(dbName).Collection(common.CONST_MO_PROJECT_NOTIFICATIONS)
	}

	timestamp := time.Now()
	var outputStruct common.SaveNotificationStruct = common.SaveNotificationStruct{userId, notificationType, subject, message, link, timestamp.Format("2006-01-02 15:04:05"), read}

	_, err := notificationCollection.InsertOne(context.TODO(), outputStruct)

	if err != nil {
		fmt.Println(err.Error())
		msg = err.Error()
	} else {
		status = true
		msg = "Success"
	}

	fmt.Println("insert method")
	fmt.Println(status, msg)

	result = Newtest{msg, status}

	return status, msg, result

}

// Fetch project details from database
func FetchNotification() (status bool, msg string, notificationDetails common.FetchNotificationStruct) {

	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	notifCollection := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_NOTIFICATIONS)
	err := notifCollection.FindOne(context.TODO(), bson.M{}).Decode(&notificationDetails)

	// err := fetchNotification.Find(context.TODO(),bson.M{"_id":"erewrwer"})

	if err != nil {
		msg = err.Error()
	} else {
		status = true
		msg = "Success"
	}

	fmt.Println("fetch method")
	fmt.Println(status, msg, notificationDetails)

	fmt.Println("before return")
	return status, msg, notificationDetails
	
}
