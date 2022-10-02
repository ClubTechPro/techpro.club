package projects

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/users"
)

type FinalOutStruct struct {
	ProgrammingLanguages map[string]string               `json:"programmingLanguages"`
	AlliedServices       map[string]string               `json:"alliedServices"`
	ProjectType          map[string]string               `json:"projectType"`
	Contributors         map[string]string               `json:"contributors"`
	UserNameImage        common.UsernameImageStruct      `json:"userNameImage"`
	NotificaitonsCount   int64                           `json:"notificationsCount"`
	NotificationsList    []common.MainNotificationStruct `json:"nofiticationsList"`
	PageDetails common.PageDetails `json:"pageDetails"`
}

func ProjectCreate(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/projects/create" {
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

	if r.Method == "GET" {

		baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title: "New Project"}

		output := FinalOutStruct{
			common.ProgrammingLanguages,
			common.AlliedServices,
			common.ProjectType,
			common.Contributors,
			userNameImage,
			notificationsCount,
			notificationsList,
			pageDetails,
		}

		tmpl, err := template.New("").ParseFiles("templates/app/common/base.gohtml", "templates/app/common/projectmenu.gohtml", "templates/app/projects/projectcreate.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			tmpl.ExecuteTemplate(w, "base", output)
		}

	} else {
		errParse := r.ParseForm()
		if errParse != nil {
			fmt.Println(errParse.Error())
		} else {
			projectName := r.Form.Get("projectName")
			repoLink := r.Form.Get("repoLink")
			projectDescription := r.Form.Get("projectDescription")
			language := r.Form["language"]
			otherLanguages := r.Form.Get("otherLanguages")
			allied := r.Form["allied"]
			projectType := r.Form["pType"]
			contributorCount := r.Form.Get("contributorCount")
			documentation := r.Form.Get("documentation")
			public := r.Form.Get("public")
			company := r.Form.Get("company")
			companyName := r.Form.Get("companyName")
			funded := r.Form.Get("funded")
			submit := r.Form.Get("submit")

			otherLanguagesSplit := strings.Split(otherLanguages, ",")

			timeNow := time.Now()
			dt := timeNow.Format(time.UnixDate)
			var result common.SaveProjectStruct

			if submit == "Save as draft" {
				result = common.SaveProjectStruct{UserID: userID, ProjectName: projectName, ProjectDescription: projectDescription, RepoLink: repoLink, Languages: language, OtherLanguages: otherLanguagesSplit, Allied: allied, ProjectType: projectType, ContributorCount: contributorCount, Documentation: documentation, Public: public, Company: company, CompanyName: companyName, Funded: funded, CreatedDate: dt, PublishedDate: "", ClosedDate: "", IsActive: common.CONST_INACTIVE, ReactionsCount: 0}
				saveProject(w, r, result)
			} else {
				result = common.SaveProjectStruct{UserID: userID, ProjectName: projectName, ProjectDescription: projectDescription, RepoLink: repoLink, Languages: language, OtherLanguages: otherLanguagesSplit, Allied: allied, ProjectType: projectType, ContributorCount: contributorCount, Documentation: documentation, Public: public, Company: company, CompanyName: companyName, Funded: funded, CreatedDate: dt, PublishedDate: dt, ClosedDate: "", IsActive: common.CONST_ACTIVE, ReactionsCount: 0}
				_, _, projectID := saveProject(w, r, result)
				go createProjectNotifications(projectID, result)
			}

			http.Redirect(w, r, "/projects/thankyou", http.StatusSeeOther)
		}
	}
}

func saveProject(w http.ResponseWriter, r *http.Request, newProjectStruct common.SaveProjectStruct) (status bool, msg string, projectId primitive.ObjectID) {
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveProject := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)

	projectDetails, err := saveProject.InsertOne(context.TODO(), newProjectStruct)

	if err != nil {
		msg = err.Error()
		projectId = primitive.NilObjectID
	} else {
		projectId = projectDetails.InsertedID.(primitive.ObjectID)
		status = true
		msg = "Success"
	}

	return status, msg, projectId
}

// Send notifications
func createProjectNotifications(projectID primitive.ObjectID, projectParams common.SaveProjectStruct) {

	var preferenceStruct common.FetchContributorPreferencesStruct
	var newNotification common.MainNotificationStruct
	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	sendNotifications := client.Database(dbName).Collection(common.CONST_MO_NOTIFICATIONS)
	contributorPreferences := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_PREFERENCES)

	var orConditionsLanguages []bson.M = []bson.M{
		{"languages": projectParams.Languages},
		{"otherlanguages": projectParams.OtherLanguages},
		{"allied": projectParams.Allied},
		{"projecttype": projectParams.ProjectType},
	}

	var contributorsArray []string
	var orConditionsContributors []bson.M
	if projectParams.ContributorCount == common.CONTRIBUTOR_COUNT[0] {
		contributorsArray = append(contributorsArray, common.CONTRIBUTOR_COUNT[2], common.CONTRIBUTOR_COUNT[1], common.CONTRIBUTOR_COUNT[0])
	} else if projectParams.ContributorCount == common.CONTRIBUTOR_COUNT[1] {
		contributorsArray = append(contributorsArray, common.CONTRIBUTOR_COUNT[1], common.CONTRIBUTOR_COUNT[0])
	} else {
		contributorsArray = append(contributorsArray, common.CONTRIBUTOR_COUNT[2])
	}
	orConditionsContributors = append(orConditionsContributors, bson.M{"contributorcount": bson.M{"$in": contributorsArray}})

	andConditions := []bson.M{
		{"$or": orConditionsLanguages},
		{"$or": orConditionsContributors},
	}

	fetchPreferences, errPreferences := contributorPreferences.Find(context.TODO(), bson.M{"userid": bson.M{"$ne": projectParams.UserID}, "$and": andConditions})
	if errPreferences != nil {
		fmt.Println(errPreferences.Error())
	} else {
		for fetchPreferences.Next(context.TODO()) {
			fetchPreferences.Decode(&preferenceStruct)

			newNotification.Link = "/projects/view?projectid=" + projectID.Hex()
			newNotification.NotificationType = common.NOTIFICATION_TYPES[0]
			newNotification.Subject = projectParams.ProjectName
			newNotification.Message = projectParams.ProjectDescription
			newNotification.CreatedDate = projectParams.CreatedDate
			newNotification.Read = false

			sendNotifications.UpdateOne(context.TODO(), bson.M{"userid": preferenceStruct.UserID}, bson.M{"$inc": bson.M{"unreadnotifications": 1}, "$push": bson.M{"notificationslist": newNotification}})
		}
	}

}
