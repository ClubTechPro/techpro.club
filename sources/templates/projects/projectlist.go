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

type ViewProjectStruct struct{
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
	IsActive bool `json:"isActive"`
}

func ProjectList(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/projects/list" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	sessionOk, userID := users.ValidateSession(w, r)
	if(!sessionOk){
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}


	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchProject := client.Database(dbName).Collection(common.CONST_PR_PROJECTS)
	projectsList, err := fetchProject.Find(context.TODO(),  bson.M{"userid": userID})

	var results []ViewProjectStruct

	if err != nil {
		fmt.Println(err)
	} else {
		for projectsList.Next(context.TODO()){
			var elem ViewProjectStruct
			errDecode := projectsList.Decode(&elem)

			if errDecode != nil {
				fmt.Println(errDecode.Error())
			} else {
				results = append(results, elem)
			}
		}
	}
	

	fmt.Println(results)
	tmpl := template.Must(template.ParseFiles("templates/app/projects/projectlist.html"))
	tmpl.Execute(w, results) 
}