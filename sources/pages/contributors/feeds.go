package contributors

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
type FinalFeedsOutputStruct struct{
	Projects []common.FeedStruct `json:"projects"`
	UserNameImage common.UsernameImageStruct `json:"usernameImage"`
	MyBookmarks []primitive.ObjectID `json:"myBookmarks"`
	MyReactions []primitive.ObjectID `json:"myReactions"`
	NotificaitonsCount int64 `json:"notificationsCount"`
	NotificationsList []common.MainNotificationStruct `json:"nofiticationsList"`
}

func Feeds(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/contributors/feeds" {
        pages.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	// Fetch notificaitons
	_, _, notificationsCount, notificationsList := pages.NotificationsCountAndTopFive(userID)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	var functions = template.FuncMap{
		"objectIdToString" : pages.ObjectIDToString,
		"containsObjectId" : pages.ContainsObjectID,
	}
	
	// TEST CONDITIONS
	// This has to come from the actual frontend
	pageid := int64(0)
	tags := []string{}
	keyword := ""

	_, _, results := filterActiveProjects(pageid, tags, keyword)
	_, _, bookmarks, reactions := pages.FetchMyBookmarksAndReactions(userID)

	output := FinalFeedsOutputStruct{results, userNameImage, bookmarks, reactions, notificationsCount, notificationsList}

	tmpl, err := template.New("").Funcs(functions).ParseFiles("templates/app/contributors/feeds.gohtml", "templates/app/contributors/common/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", output) 
	}

}


// Filter all active projects from the database according to filters
func filterActiveProjects(pageid int64, tags []string, keyword string)(status bool, msg string, results []common.FeedStruct){


	status = false
	msg = ""

	var orConditions []bson.M
	var finalConditions []bson.M
	resultsPerPage := int64(20)

	status, msg,  client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchProjects := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)

	finalConditions = append(finalConditions, bson.M{"isactive": bson.M{"$eq": common.CONST_ACTIVE}})

	if len(keyword) > 0 {
		finalConditions = append(finalConditions, bson.M{"projectname" : bson.M{"$regex": keyword}})
	}

	// Filter where conditions
	if len(tags) > 0 {
		orConditions = append(orConditions, bson.M{"languages": bson.M{"$in": tags}})
		orConditions = append(orConditions, bson.M{"otherlanguages": bson.M{"$in": tags}})
		orConditions = append(orConditions, bson.M{"allied": bson.M{"$in": tags}})

		finalConditions = append(finalConditions, bson.M{"$or" : orConditions})
	}
	

	aggCondition := bson.M{"$match": bson.M{"$and" : finalConditions}}


	// Filter joins
	aggLookup := bson.M{"$lookup": bson.M{
		"from":         common.CONST_MO_USERS,    // the collection name
		"localField":   "userid", 	      		  // the field on the child struct
		"foreignField": "_id",       		  	  // the field on the parent struct
		"as":           "userdetails",    		  // the field to populate into
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
		"reactionscount": 1,
		"public" : 1,
		"userdetails" : bson.M{ "_id" : 1, "name": 1, "imagelink" :1},
	}}

	aggSkip := bson.M{"$skip": (pageid * resultsPerPage)}
    aggLimit := bson.M{"$limit": resultsPerPage}

	projectsList, err := fetchProjects.Aggregate(context.TODO(), []bson.M{aggCondition, aggLookup, aggProjections, aggSkip, aggLimit})

	if err != nil {
		msg = err.Error()
	} else {
		for projectsList.Next(context.TODO()){
			var elem common.FeedStruct
			errDecode := projectsList.Decode(&elem)

			if errDecode != nil {
				msg = errDecode.Error()
			} else {
				results = append(results, elem)
				status = true
				msg = "Success"
			}
		}

	}

	return status, msg, results
}
