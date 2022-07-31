package pages

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
	"techpro.club/sources/users"
)

// Fetch user name and image from saved browser cookies
func FetchUsernameImage(w http.ResponseWriter, r *http.Request) (status bool, msg, userName, image string) {
	status = false
	msg = ""
	
	// user name cookie
	userNameCookie, err := r.Cookie(common.CONST_USER_NAME)

	if err != nil {
		status = false
		msg = err.Error()
		userName = ""
	} else {
		status = true
		msg = "Success"
		userName = userNameCookie.Value
	}

	// user image cookie
	imageCookie, err := r.Cookie(common.CONST_USER_IMAGE)

	if err != nil {
		status = false
		msg = err.Error()
		image = ""
	} else {
		status = true
		msg = "Success"
		image = imageCookie.Value
	}

	return status, msg, userName, image
}


// Check if a string exists in a slice.
func Contains(s []string, e string) (status bool) {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

// Convert slice of strings to csv string
func SliceToCsv(s []string) (csv string){
	csv = strings.Join(s, ",")
	return csv
}

// Fetch project details from database
func FetchProjectDetails(projectID string, userID primitive.ObjectID) (status bool, msg string, projectDetails common.FetchProjectStruct){

	status = false
	msg = ""

	if(projectID != ""){

		projectIdHex, err := primitive.ObjectIDFromHex(projectID)

		if err != nil {
			msg = err.Error()
		} else {
			_, _, client := common.Mongoconnect()
			defer client.Disconnect(context.TODO())

			dbName := common.GetMoDb()
			fetchProject := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)
			err := fetchProject.FindOne(context.TODO(),  bson.M{"userid": userID, "_id": projectIdHex}).Decode(&projectDetails)

			if err != nil {
				msg = err.Error()
			}  else {
				status = true
				msg = "Success"
			}
		}
	} else {
		msg = "Project ID is empty"
	}

	return status, msg, projectDetails
}

// Find total unread notifications for a user from database
func NotificationsCount(userID primitive.ObjectID)(status bool, msg string, count int64){
	status = true
	msg = "Success"
	count = 0

	status, msg, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	countNotifications := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_NOTIFICATIONS)
	count, errCount := countNotifications.CountDocuments(context.TODO(), bson.M{"read": false, "userid" : userID})

	if errCount != nil{
		status = false
		msg = errCount.Error()
	} else {
		status = true
		msg = "Success"
	}

	return status, msg, count
}

// Manage reaction to a project
func ManageReactions(w http.ResponseWriter, r *http.Request){

	msg := ""
	status := false

	type InputStruct struct{
		ProjectId primitive.ObjectID `json:"projectid"`
	}

	var inputJSON InputStruct

	_, userID := users.ValidateDbSession(w, r)

	readData, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Println("Err1", errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println("Err2", errParse)

	} 

	if(inputJSON.ProjectId != primitive.NilObjectID){

		_, _, client := common.Mongoconnect()
		defer client.Disconnect(context.TODO())

		dbName := common.GetMoDb()

		var projectIdList []primitive.ObjectID
		projectIdList = append(projectIdList, inputJSON.ProjectId)

		fetchProjectReactions := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)

		fetchUserProjectReactions := client.Database(dbName).Collection(common.CONST_MO_USER_PROJECT_REACTIONS)
		result, err := fetchUserProjectReactions.CountDocuments(context.TODO(), bson.M{"userid" : userID})

		if err != nil {
			msg = err.Error()
		}  else {

			// If it contains the project, then delete it
			// Else insert it
			if(result > 0){

				resultCountProjects, errCountProjects := fetchUserProjectReactions.CountDocuments(context.TODO(), bson.M{"userid" : userID, "projectids" : bson.M{"$in" : projectIdList}})

				if errCountProjects != nil {
					msg = errCountProjects.Error()
				} else {

					if (resultCountProjects > 0){
						_, errProjectReactions := fetchProjectReactions.UpdateOne(context.TODO(), bson.M{"_id": inputJSON.ProjectId}, bson.M{"$inc" : bson.M{"reactionscount" : -1}})
						_, err := fetchUserProjectReactions.UpdateOne(context.TODO(), bson.M{"userid": userID}, bson.M{"$pull" : bson.M{"projectids" : inputJSON.ProjectId}})
						if err != nil || errProjectReactions != nil {
							msg = err.Error() + ". " + errProjectReactions.Error()
						}  else {
							status = true
							msg = "Success"
						}
					} else {
						_, errProjectReactions := fetchProjectReactions.UpdateOne(context.TODO(), bson.M{"_id": inputJSON.ProjectId}, bson.M{"$inc" : bson.M{"reactionscount" : 1}})
						_, err := fetchUserProjectReactions.UpdateOne(context.TODO(), bson.M{"userid": userID}, bson.M{"$push" : bson.M{"projectids" : inputJSON.ProjectId}})
						if err != nil  || errProjectReactions != nil {
							msg = err.Error() + ". " + errProjectReactions.Error()
						}  else {
							status = true
							msg = "Success"
						}
					}
					
				}
				
			} else {

				var userProjectReactions common.SaveUserProjectReactionStruct
				userProjectReactions.UserId = userID
				userProjectReactions.ProjectIds = projectIdList

				_, errUserProjectReactions := fetchUserProjectReactions.InsertOne(context.TODO(), userProjectReactions)
				if  errUserProjectReactions != nil {
					msg =  errUserProjectReactions.Error()
				}  else {
					status = true
					msg = "Success"
				}
			}	
		}
				
	}

	output := common.JsonOutput{
		Status: status,
		Msg: msg,
	}

	out, _ := json.Marshal(output)
	w.Write(out)
}

// Manage bookmarks to a project
func ManageBookmarks(w http.ResponseWriter, r *http.Request){

	msg := ""
	status := false

	type InputStruct struct{
		ProjectId primitive.ObjectID `json:"projectid"`
	}

	var inputJSON InputStruct

	_, userID := users.ValidateDbSession(w, r)

	readData, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		log.Println("Err1", errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println("Err2", errParse)

	} 

	if(inputJSON.ProjectId != primitive.NilObjectID){

		_, _, client := common.Mongoconnect()
		defer client.Disconnect(context.TODO())

		dbName := common.GetMoDb()

		var projectIdList []primitive.ObjectID
		projectIdList = append(projectIdList, inputJSON.ProjectId)

		fetchUserProjectBookmarks := client.Database(dbName).Collection(common.CONST_MO_BOOKMARKS)
		result, err := fetchUserProjectBookmarks.CountDocuments(context.TODO(), bson.M{"userid" : userID})

		if err != nil {
			msg = err.Error()
		}  else {

			// If it contains the project, then delete it
			// Else insert it
			if(result > 0){

				resultCountProjects, errCountProjects := fetchUserProjectBookmarks.CountDocuments(context.TODO(), bson.M{"userid" : userID, "projectids" : bson.M{"$in" : projectIdList}})

				if errCountProjects != nil {
					msg = errCountProjects.Error()
				} else {

					if (resultCountProjects > 0){
						_, err := fetchUserProjectBookmarks.UpdateOne(context.TODO(), bson.M{"userid": userID}, bson.M{"$pull" : bson.M{"projectids" : inputJSON.ProjectId}})
						if err != nil {
							msg = err.Error()
						}  else {
							status = true
							msg = "Success"
						}
					} else {
						_, err := fetchUserProjectBookmarks.UpdateOne(context.TODO(), bson.M{"userid": userID}, bson.M{"$push" : bson.M{"projectids" : inputJSON.ProjectId}})
						if err != nil {
							msg = err.Error()
						}  else {
							status = true
							msg = "Success"
						}
					}
					
				}
				
			} else {
				var userProjectReactions common.SaveUserProjectReactionStruct
				userProjectReactions.UserId = userID
				userProjectReactions.ProjectIds = projectIdList

				_, err := fetchUserProjectBookmarks.InsertOne(context.TODO(), userProjectReactions)
				if err != nil {
					msg = err.Error()
				}  else {
					status = true
					msg = "Success"
				}
			}	
		}
				
	}

	output := common.JsonOutput{
		Status: status,
		Msg: msg,
	}

	out, _ := json.Marshal(output)
	w.Write(out)
}

// Convert primitive.ObjectID to string
func ObjectIDToString(Id primitive.ObjectID)(idString string){

	return Id.Hex()
}

// Convert string to primitive.ObjectID 
func StringToObjectId(Id string)(idObject primitive.ObjectID){
	idObject, err := primitive.ObjectIDFromHex(Id)

	if err != nil {
		fmt.Println(err.Error())
	}

	return idObject
}