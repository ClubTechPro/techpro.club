package projects

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"techpro.club/sources/common"
	"techpro.club/sources/users"
)

type NewProjectStruct struct{
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

func ProjectCreate(w http.ResponseWriter, r *http.Request){

	sessionOk, userID := users.ValidateSession(w, r)
	if(!sessionOk){
		fmt.Println(sessionOk, userID)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	
	if r.Method == "GET"{
		tmpl := template.Must(template.ParseFiles("templates/app/projects/projectcreate.html"))
		tmpl.Execute(w, nil) 
	} else {
		errParse := r.ParseForm()
		if errParse != nil {
			fmt.Println(errParse.Error())
		} else {
			projectName := r.Form.Get("projectName")
			repoLink := r.Form.Get("repoLink")
			projectDescription := r.Form.Get("projectDescription")
			language := r.Form["language"]
			otherLanguages := r.Form["otherLanguages"]
			allied := r.Form["allied"]
			projectType :=  r.Form["pType"]
			contributorCount := r.Form.Get("contributorCount")
			documentation := r.Form.Get("documentation")
			public := r.Form.Get("public")
			company := r.Form.Get("company")
			companyName := r.Form.Get("companyName")
			funded := r.Form.Get("funded")
			submit := r.Form.Get("submit")

			time := time.Now()
			var result NewProjectStruct

			if submit == "Save as draft" {
				result = NewProjectStruct{userID, projectName, repoLink, projectDescription, language, otherLanguages, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, time.String(), "", "", false}
				saveProject(w, r, result)
			} else {
				result = NewProjectStruct{userID, projectName, repoLink, projectDescription, language, otherLanguages, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, time.String(), time.String(), "", true}
				saveProject(w, r, result)
			}	
		}
	}
	
}


func saveProject(w http.ResponseWriter, r *http.Request, newProjectStruct NewProjectStruct){
	client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveProject := client.Database(dbName).Collection(common.CONST_PR_PROJECTS)

	_, err := saveProject.InsertOne(context.TODO(), newProjectStruct)

	if err != nil {
		fmt.Println(err)
	}
	
	

	http.Redirect(w, r, "/projects/thankyou", http.StatusSeeOther)
}