package projects

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)

type ProjectStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID string `json:"userId"`
	ProjectName string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
	RepoLink string `json:"repoLink"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	ProjectType []string `json:"projectType"`
	ContributorCount string `json:"contributorCount"`
	Documentation string `json:"documentation"`
	Public string `json:"public"`
	Company string `json:"company"`
	CompanyName string `json:"companyName"`
	Funded string `json:"funded"`
	CreatedDate string `json:"createdDate"`
	PublishedDate string `json:"publishedDate"`
	ClosedDate string `json:"closedDate"`
	IsActive int `json:"isActive"`
}

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

	var results []ProjectStruct

	if err != nil {
		fmt.Println(err)
	} else {
		for projectsList.Next(context.TODO()){
			var elem ProjectStruct
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