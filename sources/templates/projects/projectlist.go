package projects

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"techpro.club/sources/common"
	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)


func ProjectList(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/projects/list" {
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


	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchProject := client.Database(dbName).Collection(common.CONST_PR_PROJECTS)
	projectsList, err := fetchProject.Find(context.TODO(),  bson.M{"userid": userID})

	var results []common.FetchProjectStruct

	if err != nil {
		fmt.Println(err)
	} else {
		for projectsList.Next(context.TODO()){
			var elem common.FetchProjectStruct
			errDecode := projectsList.Decode(&elem)

			if errDecode != nil {
				fmt.Println(errDecode.Error())
			} else {
				results = append(results, elem)
			}
		}
	}
	

	tmpl, err := template.New("").ParseFiles("templates/app/projects/projectlist.gohtml", "templates/app/projects/common/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "projectbase", results) 
	}
}