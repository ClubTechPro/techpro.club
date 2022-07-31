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

type FinalReactionsOutputStruct struct{
	Projects []common.FeedStruct `json:"projects"`
	UserNameImage common.UsernameImageStruct `json:"usernameImage"`
	MyBookmarks []primitive.ObjectID `json:"myBookmarks"`
}

// Fetched reacted projects
func FetchReactions(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contributors/reactions" {
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

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	var functions = template.FuncMap{
		"objectIdToString" : pages.ObjectIDToString,
		"containsObjectId" : pages.ContainsObjectID,
	}

	// constants for check
	pageid := 0

	// Fetch all reacted projects
	// Also fetch project details where the user reacted
	_, _, results := fetchReactedProjectsList(int64(pageid), userID)
	_, _, bookmarks, _ := pages.FetchMyBookmarksAndReactions(userID)
	

	output := FinalReactionsOutputStruct{results, userNameImage, bookmarks}

	tmpl, err := template.New("").Funcs(functions).ParseFiles("templates/app/contributors/common/base.gohtml", "templates/app/contributors/reactions.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", output) 
	}
}

// Fetch all reacted projects
func filterReactedProjects(pageid int64, userID primitive.ObjectID)(status bool, msg string, results []primitive.ObjectID){


	status = false
	msg = ""
	var out common.FetchUserProjectReactionStruct

	status, msg,  client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchReactedProjects := client.Database(dbName).Collection(common.CONST_MO_USER_PROJECT_REACTIONS)

	// Fetch all reacted projects against a user
	reactedProjectsResults, err:= fetchReactedProjects.Find(context.TODO(), bson.M{"userid": userID})

	if err != nil{
		msg = err.Error()
	} else {
		for reactedProjectsResults.Next(context.TODO()){
			errDecode := reactedProjectsResults.Decode(&out)

			if errDecode != nil {
				msg = errDecode.Error()
			} else {
				results = out.ProjectIds
			}
		}
	}

	return status, msg, results
}


// Filter all active projects from the database according to users reactions
func fetchReactedProjectsList(pageid int64, userID primitive.ObjectID)(status bool, msg string, results []common.FeedStruct){


	status = false
	msg = ""

	var finalConditions []bson.M
	resultsPerPage := int64(20)

	status, msg,  client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchProjects := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)

	finalConditions = append(finalConditions, bson.M{"isactive": bson.M{"$eq": common.CONST_ACTIVE}})

	status, errMsg, projectIds := filterReactedProjects(pageid, userID)

	if !status{
		msg = errMsg
	} else {
		if len(projectIds) <= 0{
			msg = "No projects found"
			results = []common.FeedStruct{}			
		} else {
			finalConditions = append(finalConditions, bson.M{"_id" : bson.M{"$in": projectIds}})
			
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
				"public" : 1,
				"reactionscount": 1,
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
		}
	}

	return status, msg, results
}

