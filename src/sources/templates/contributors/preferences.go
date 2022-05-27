package contributors

import (
	"config"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"sources/common"
)


type ContributorPreferencesStruct struct{
	Languages []string `json:"languages"`
	NotificationFrequency string `json:"notificationFrequency"`
}

func ContributorPreferences(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("../templates/app/contributors/preferences.html"))
    tmpl.Execute(w, nil)

	if r.Method == "POST"{

		r.ParseForm()
		languages := r.Form["language"]
		notificationFrequency := r.Form.Get("emailFrequency")

		result := ContributorPreferencesStruct{languages, notificationFrequency}

		client := config.Mongoconnect()
		defer client.Disconnect(context.TODO())

		dbName := common.GetMoDb()
		saveContributorPreference := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_PREFERENCES)

		_, err := saveContributorPreference.InsertOne(context.TODO(), result)

		if err != nil {
			fmt.Println(err)
		}

	}
}