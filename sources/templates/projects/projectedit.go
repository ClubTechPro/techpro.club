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
		"contains" : Contains,
		"sliceToCsv" : SliceToCsv,
	}

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, userName, image := templates.FetchUsernameImage(w, r)

	if(!status){
		log.Println("Error fetching user name and image from cookies")
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	if r.Method == "GET"{

		projectID := r.URL.Query().Get("projectid")

		result := FetchProjectDetails(projectID, userID)

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
		}
	}
}

func FetchProjectDetails(projectID string, userID primitive.ObjectID) (projectDetails common.FetchProjectStruct){
	if(projectID != ""){

		projectIdHex, err := primitive.ObjectIDFromHex(projectID)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			client, _ := common.Mongoconnect()
			defer client.Disconnect(context.TODO())

			dbName := common.GetMoDb()
			fetchProject := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)
			err := fetchProject.FindOne(context.TODO(),  bson.M{"userid": userID, "_id": projectIdHex}).Decode(&projectDetails)

			if err != nil {
				fmt.Println(err)
			} 
		}
	}

	return projectDetails
}

func updateProject(w http.ResponseWriter, r *http.Request, projectID string, newProjectStruct common.SaveProjectStruct){
	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveProject := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)

	projectIdHex, _ := primitive.ObjectIDFromHex(projectID)
	_, err := saveProject.UpdateOne(context.TODO(), bson.M{"_id": projectIdHex}, bson.M{"$set": newProjectStruct})

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/projects/thankyou", http.StatusSeeOther)
}

// Check if a string exists in a slice.
func Contains(s []string, e string) (status bool) {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

// Convert slice of strings to csv string
func SliceToCsv(s []string) (csv string){
	csv = strings.Join(s, ",")
	return csv
}