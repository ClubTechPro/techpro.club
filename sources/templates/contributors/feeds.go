package contributors

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"techpro.club/sources/common"
	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)

func Feeds(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/contributors/feeds" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	sessionOk, _ := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	
	fetchActiveProjects(0)

	tmpl, err := template.New("").ParseFiles("templates/app/contributors/feeds.gohtml", "templates/app/contributors/common/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", nil) 
	}

}

func fetchActiveProjects(pageid int64)(results []common.FetchProjectStruct){

	resultsPerPage := int64(20)

	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchProjects := client.Database(dbName).Collection(common.CONST_PR_PROJECTS)
	projectsList, err := fetchProjects.Find(context.TODO(),  bson.M{"isactive": common.CONST_ACTIVE}, options.Find().SetLimit(resultsPerPage).SetSkip(pageid*resultsPerPage))
	
	if err != nil {
		fmt.Println(err.Error())
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

	return results
}