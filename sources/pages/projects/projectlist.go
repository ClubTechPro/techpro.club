package projects

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/users"
)

type FinalProjectListOutStruct struct{
	ProjectsList []common.FetchProjectStruct `json:"projectsList"`
	UserNameImage common.UsernameImageStruct `json:"userNameImage"`
}

func ProjectList(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/projects/list" {
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

	var finalOutStruct FinalProjectListOutStruct
	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}
	
	_, _, results := listProjects(w, r, userID)
	finalOutStruct = FinalProjectListOutStruct{results, userNameImage}
	

	tmpl, err := template.New("").ParseFiles("templates/app/projects/projectlist.gohtml", "templates/app/projects/common/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "projectbase", finalOutStruct) 
	}
}


// List projects
func listProjects(w http.ResponseWriter, r *http.Request, userID primitive.ObjectID)(status bool, msg string, results []common.FetchProjectStruct){
	
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchProject := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)
	projectsList, err := fetchProject.Find(context.TODO(),  bson.M{"userid": userID})


	if err != nil {
		msg = err.Error()
	} else {
		for projectsList.Next(context.TODO()){
			var elem common.FetchProjectStruct
			errDecode := projectsList.Decode(&elem)

			if errDecode != nil {
				msg = errDecode.Error()
				status = false
			} else {
				status = true
				msg = "Success"
				results = append(results, elem)
			}
		}
	}

	return status, msg, results
}