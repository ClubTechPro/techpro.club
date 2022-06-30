package contributors

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"techpro.club/sources/common"
	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)
type FinalFeedsOutputStruct struct{
	Projects []common.FeedStruct `json:"projects"`
	UserNameImage common.UsernameImageStruct `json:"usernameImage"`
}

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
	

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, userName, image := templates.FetchUsernameImage(w, r)

	if(!status){
		log.Println("Error fetching user name and image from cookies")
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}
	
	results := fetchActiveProjects(0)

	output := FinalFeedsOutputStruct{results, userNameImage}

	tmpl, err := template.New("").ParseFiles("templates/app/contributors/feeds.gohtml", "templates/app/contributors/common/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", output) 
	}

}

func fetchActiveProjects(pageid int64)(results []common.FeedStruct){

	resultsPerPage := int64(20)

	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchProjects := client.Database(dbName).Collection(common.CONST_PR_PROJECTS)

	aggLookup := bson.M{"$lookup": bson.M{
		"from":         common.CONST_MO_USERS,    // the collection name
		"localField":   "projects.userid", 	      // the field on the child struct
		"foreignField": "users._id",       		  // the field on the parent struct
		"as":           "userdetails",    			  // the field to populate into
	}}

	// Set projections
	aggProjections := bson.M{"$project": bson.M{ 
		"_id": 1, "projectname" : 1, 
		"projectdescription" : 1, 
		"repolink": 1, 
		"languages": 1, 
		"otherlanguages": 1, 
		"allied": 1, 
		"company" : 1, 
		"companyname": 1, 
		"createddate": 1,
		"userdetails" : bson.M{ "_id" : 1, "name": 1, "imagelink" :1},
	}}

	aggCondition := bson.M{"$match": bson.M{"isactive": bson.M{"$eq": common.CONST_ACTIVE}}}

	aggSkip := bson.M{"$skip": (pageid * resultsPerPage)}
    aggLimit := bson.M{"$limit": resultsPerPage}

	projectsList, err := fetchProjects.Aggregate(context.TODO(), []bson.M{aggCondition, aggLookup, aggProjections, aggSkip, aggLimit})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for projectsList.Next(context.TODO()){
			var elem common.FeedStruct
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