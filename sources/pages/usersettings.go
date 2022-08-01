package pages

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"techpro.club/sources/common"
	"techpro.club/sources/users"
)

type UserSettingsStruct struct {
	UserProfile   common.FetchUserStruct     `json:"userprofile"`
	UserNameImage common.UsernameImageStruct `json:"usernameImage"`
	NotificaitonsCount int64 `json:"notificationsCount"`
	NotificationsList []common.MainNotificationStruct `json:"nofiticationsList"`
}

// Display user settings
func UserSettings(w http.ResponseWriter, r *http.Request) {

	// Check if the path is ok
	// Check if the session is ok

	// Fetch name and image from cookies
	// Fetch email, _id, location,repourl from database
	// Display the user's profile
	// Display the user's settings

	if r.URL.Path != "/users/settings" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if !sessionOk {

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

	if !status {
		log.Println(msg)
	} else {
		userNameImage = common.UsernameImageStruct{userName, image}
	}

	_, _, userprofile := fetchUserProfile(userID)

	userSettingData := UserSettingsStruct{userprofile, userNameImage, notificationsCount, notificationsList}

	tmpl, err := template.New("").ParseFiles("templates/app/contributors/settings.gohtml", "templates/app/contributors/common/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		tmpl.ExecuteTemplate(w, "contributorbase", userSettingData)
	}
}

// Manage reaction to a project
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	status := false
	msg := ""

	if r.URL.Path != "/api/deleteuser" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if !sessionOk {

		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	type Test struct {
		UserEmail string `json:"userEmail"`
	}

	var inputJSON Test

	readData, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Println("Err1", errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println("Err2", errParse)

	}

	_, _, userprofile := fetchUserProfile(userID)

	if strings.ToLower(inputJSON.UserEmail) == strings.ToLower(userprofile.Email) {

		_, _, client := common.Mongoconnect()
		defer client.Disconnect(context.TODO())

		dbName := common.GetMoDb()

		projectCollection := client.Database(dbName).Collection(common.CONST_MO_NOTIFICATIONS)
		_, err1 := projectCollection.DeleteMany(context.TODO(), bson.M{"userid": userID})

		userSessionCollection := client.Database(dbName).Collection(common.CONST_MO_USER_SESSIONS)
		_, err2 := userSessionCollection.DeleteMany(context.TODO(), bson.M{"userid": userID})

		userCollection := client.Database(dbName).Collection(common.CONST_MO_USERS)
		_, err3 := userCollection.DeleteOne(context.TODO(), bson.M{"_id": userID})

		projectsCollection := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)
		_, err4 := projectsCollection.DeleteOne(context.TODO(), bson.M{"userid": userID})

		socialsCollection := client.Database(dbName).Collection(common.CONST_MO_SOCIALS)
		_, err5 := socialsCollection.DeleteOne(context.TODO(), bson.M{"userid": userID})

		if err1 != nil && err2 != nil && err3 != nil && err4 != nil && err5 != nil{
			msg = err1.Error()
		} else {
			status = true
			msg = "Success"
		}

		fmt.Println(status, msg)

	}
	resp := make(map[string]string)
	resp["message"] = "User account remove successfully"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
