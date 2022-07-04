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
	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)

type FinalProjectOutStruct struct{
	ProgrammingLanguages map[string]string `json:"programmingLanguages"`
	AlliedServices map[string]string `json:"alliedServices"`
	ProjectType map[string]string `json:"projectType"`
	Contributors map[string]string `json:"contributors"`
	ProjectStruct common.FetchProjectStruct `json:"projectStruct"`
	UserNameImage common.UsernameImageStruct `json:"userNameImage"`
}

func ProjectEdit(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/projects/edit" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
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
		"contains" : templates.Contains,
		"sliceToCsv" : templates.SliceToCsv,
	}

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := templates.FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	if r.Method == "GET"{

		projectID := r.URL.Query().Get("projectid")

		_, _, result := templates.FetchProjectDetails(projectID, userID)

		constantLists := FinalProjectOutStruct{
			common.ProgrammingLanguages,
			common.AlliedServices,
			common.ProjectType,
			common.Contributors,
			result,
			userNameImage,
		}
		

		tmpl, err := template.New("").Funcs(functions).ParseFiles("templates/app/projects/projectedit.gohtml", "templates/app/projects/common/base.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "projectbase", constantLists) 
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
			var result common.SaveProjectStruct

			if submit == "Save as draft" {
				result = common.SaveProjectStruct{userID, projectName, projectDescription, repoLink, language, otherLanguagesSplit, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, dt, "", "", common.CONST_INACTIVE}
				updateProject(w, r, projectID, result)
			} else {
				result = common.SaveProjectStruct{userID, projectName, projectDescription, repoLink, language, otherLanguagesSplit, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, dt, dt, "", common.CONST_UNDER_MODERATION}
				updateProject(w, r, projectID, result)
			}
			
			http.Redirect(w, r, "/projects/thankyou", http.StatusSeeOther)
		}
	}
}

// Update project function
func updateProject(w http.ResponseWriter, r *http.Request, projectID string, newProjectStruct common.SaveProjectStruct)(status bool, msg string){
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