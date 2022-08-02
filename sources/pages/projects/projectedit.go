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

type FinalProjectOutStruct struct{
	ProgrammingLanguages map[string]string `json:"programmingLanguages"`
	AlliedServices map[string]string `json:"alliedServices"`
	ProjectType map[string]string `json:"projectType"`
	Contributors map[string]string `json:"contributors"`
	ProjectStruct common.FetchProjectStruct `json:"projectStruct"`
	UserNameImage common.UsernameImageStruct `json:"userNameImage"`
	NotificaitonsCount int64 `json:"notificationsCount"`
	NotificationsList []common.MainNotificationStruct `json:"nofiticationsList"`
}

func ProjectEdit(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/projects/edit" {
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

	var functions = template.FuncMap{
		"contains" : pages.Contains,
		"sliceToCsv" : pages.SliceToCsv,
	}

	// Fetch notificaitons
	_, _, notificationsCount, notificationsList := pages.NotificationsCountAndTopFive(userID)

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)
	

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	if r.Method == "GET"{

		projectID := r.URL.Query().Get("projectid")

		_, _, result := pages.FetchProjectDetails(projectID, userID)

		constantLists := FinalProjectOutStruct{
			common.ProgrammingLanguages,
			common.AlliedServices,
			common.ProjectType,
			common.Contributors,
			result,
			userNameImage,
			notificationsCount,
			notificationsList,
		}
		

		tmpl, err := template.New("").Funcs(functions).ParseFiles("templates/app/common/base.gohtml", "templates/app/common/projectmenu.gohtml", "templates/app/projects/projectedit.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "base", constantLists) 
		}
	} else {
		projectID := r.URL.Query().Get("projectid")

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
			projectType :=  r.Form["pType"]
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
			var result common.UpdateProjectStruct

			if submit == "Save as draft" {
				result = common.UpdateProjectStruct{userID, projectName, projectDescription, repoLink, language, otherLanguagesSplit, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, dt, "", "", common.CONST_INACTIVE}
				updateProject(w, r, projectID, result)
			} else {
				result = common.UpdateProjectStruct{userID, projectName, projectDescription, repoLink, language, otherLanguagesSplit, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, dt, dt, "", common.CONST_ACTIVE}
				updateProject(w, r, projectID, result)
			}
			
			http.Redirect(w, r, "/projects/thankyou", http.StatusSeeOther)
		}
	}
}

// Update project function
func updateProject(w http.ResponseWriter, r *http.Request, projectID string, newProjectStruct common.UpdateProjectStruct)(status bool, msg string){
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveProject := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)

	projectIdHex, _ := primitive.ObjectIDFromHex(projectID)
	_, err := saveProject.UpdateOne(context.TODO(), bson.M{"_id": projectIdHex}, bson.M{"$set": newProjectStruct})

	if err != nil {
		fmt.Println(err)
		msg = err.Error()
	} else {
		status = true
		msg = "Success"
	}

	return status, msg
}