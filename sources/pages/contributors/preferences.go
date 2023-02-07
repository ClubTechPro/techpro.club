package contributors

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/users"
)

type FinalPreferencesOutStruct struct {
	ProgrammingLanguages   map[string]string                       `json:"programmingLanguages"`
	AlliedServices         map[string]string                       `json:"alliedServices"`
	ProjectType            map[string]string                       `json:"projectType"`
	Contributors           map[string]string                       `json:"contributors"`
	ContributorPreferences common.SaveContributorPreferencesStruct `json:"contributorPreferences"`
	UserNameImage          common.UsernameImageStruct              `json:"userNameImage"`
	NotificaitonsCount     int64                                   `json:"notificationsCount"`
	NotificationsList      []common.MainNotificationStruct         `json:"nofiticationsList"`
	PageDetails            common.PageDetails                      `json:"pageDetails"`
}

func Preferences(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/contributors/preferences" {
		pages.ErrorHandler(w, r, http.StatusNotFound)
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

	var functions = template.FuncMap{
		"contains":   pages.Contains,
		"sliceToCsv": pages.SliceToCsv,
	}

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	// Fetch notificaitons
	_, _, notificationsCount, notificationsList := pages.NotificationsCountAndTopFive(userID)

	if !status {
		log.Println(msg)
	} else {
		userNameImage = common.UsernameImageStruct{Username: userName, Image: image}
	}

	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title: "Project Preferences"}

	if r.Method == "GET" {
		_, _, preferences := fetchPreferences(userID)

		finalPreferencesOutStruct := FinalPreferencesOutStruct{
			common.ProgrammingLanguages,
			common.AlliedServices,
			common.ProjectType,
			common.Contributors,
			preferences,
			userNameImage,
			notificationsCount,
			notificationsList,
			pageDetails,
		}

		tmpl, err := template.New("").Funcs(functions).ParseFiles("templates/app/common/base.gohtml", "templates/app/common/contributormenu.gohtml", "templates/app/contributors/preferences.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			tmpl.ExecuteTemplate(w, "base", finalPreferencesOutStruct)
		}
	} else {

		status, msg := savePreferences(w, r, userID)
		if !status {
			fmt.Println(msg)
		} else {
			http.Redirect(w, r, "/contributors/thankyou", http.StatusOK)
		}
	}
}

// Return contributor preferences, if already saved
func fetchPreferences(userID primitive.ObjectID) (status bool, msg string, preferences common.SaveContributorPreferencesStruct) {
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchPreferences := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_PREFERENCES)
	results, err := fetchPreferences.Find(context.TODO(), bson.M{"userid": userID}, options.Find().SetProjection(bson.M{"_id": 0}))

	if err != nil {
		msg = err.Error()
		status = false
	} else {

		for results.Next(context.TODO()) {

			errDecode := results.Decode(&preferences)

			if errDecode != nil {
				msg = errDecode.Error()
				status = false
			} else {
				msg = "Success"
				status = true
			}
		}
	}

	return status, msg, preferences
}

// Save preferences for contributor
func savePreferences(w http.ResponseWriter, r *http.Request, userID primitive.ObjectID) (status bool, msg string) {
	status = false
	msg = ""

	errParse := r.ParseForm()
	if errParse != nil {
		msg = errParse.Error()
	} else {
		languages := r.Form["language"]
		otherLanguages := r.Form.Get("otherLanguages")
		allied := r.Form["allied"]
		notificationFrequency := r.Form.Get("emailFrequency")
		projectType := r.Form["pType"]
		contributorCount := r.Form.Get("contributorCount")
		paidJob := r.Form.Get("paidJob")
		relocation := r.Form.Get("relocation")
		qualification := r.Form.Get("qualification")

		otherLanguagesSplit := strings.Split(otherLanguages, ",")

		result := common.SaveContributorPreferencesStruct{UserID: userID,
			Languages:             languages,
			OtherLanguages:        otherLanguagesSplit,
			Allied:                allied,
			ProjectType:           projectType,
			NotificationFrequency: notificationFrequency,
			ContributorCount:      contributorCount,
			PaidJob:               paidJob,
			Relocation:            relocation,
			Qualification:         qualification}

		_, _, client := common.Mongoconnect()
		defer client.Disconnect(context.TODO())

		dbName := common.GetMoDb()
		saveContributorPreference := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_PREFERENCES)

		_, err := saveContributorPreference.InsertOne(context.TODO(), result)

		if err != nil {
			msg = err.Error()
		} else {
			status = true
			msg = "Success"
		}
	}

	return status, msg
}
