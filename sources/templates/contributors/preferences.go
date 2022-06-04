package contributors

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/users"
)


type ContributorPreferencesStruct struct{
	UserID string `json:"userId"`
	Languages []string `json:"languages"`
	NotificationFrequency string `json:"notificationFrequency"`
	ProjectType []string `json:"projectType"`
	ContributorCount string `json:"contributorCount"`
	PaidJob string `json:"paidJob"`
	Relocation string `json:"relocation"`
	Qualification string `json:"qualification"`
}

func Preferences(w http.ResponseWriter, r *http.Request){
	
	sessionOk, userID := users.ValidateSession(w, r)
	if(!sessionOk){
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.Method == "GET"{
		tmpl := template.Must(template.ParseFiles("../templates/app/contributors/preferences.html"))
		tmpl.Execute(w, nil) 
	} else {
	
		errParse := r.ParseForm()
		if errParse != nil {
			log.Println(errParse.Error())
		} else {
			languages := r.Form["language"]
			notificationFrequency := r.Form.Get("emailFrequency")
			projectType := r.Form["pType"]
			contributorCount := r.Form.Get("contributorCount")
			paidJob :=  r.Form.Get("paidJob")
			relocation := r.Form.Get("relocation")
			qualification := r.Form.Get("qualification")

			result := ContributorPreferencesStruct{userID, languages, notificationFrequency, projectType, contributorCount, paidJob, relocation, qualification}

			client := common.Mongoconnect()
			defer client.Disconnect(context.TODO())
	
			dbName := common.GetMoDb()
			saveContributorPreference := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_PREFERENCES)
	
			_, err := saveContributorPreference.InsertOne(context.TODO(), result)
	
			if err != nil {
				fmt.Println(err)
			}
			
			

			http.Redirect(w, r, "/contributor/thankyou", http.StatusSeeOther)
		}
	}
}