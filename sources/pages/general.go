package pages

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// Check if a primitive.ObjectID exists in a slice.
func ContainsObjectID(o []primitive.ObjectID, e primitive.ObjectID) (status bool) {
    for _, a := range o {
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
func NotificationsCountAndTopFive(userID primitive.ObjectID)(status bool, msg string, count int64, notificationsList []common.MainNotificationStruct){
	status = false
	msg = "Failed"
	count = 0

	statusCount := false
	msgCount := ""

	statusList := false
	msgList := ""

	var notifications common.FetchNotificationStruct

	status, msg, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	countNotifications := client.Database(dbName).Collection(common.CONST_MO_NOTIFICATIONS)
	errCount := countNotifications.FindOne(context.TODO(), bson.M{"userid" : userID}).Decode(&notifications)
	if errCount != nil{
		msgCount = errCount.Error()
	} else {
		statusCount = true
		msgCount = "Success"
		count = int64(notifications.UnreadNotifications)
	}
	
	fetchNotifications, errFetch := countNotifications.Find(context.TODO(),  bson.M{"userid": userID})

	if errFetch != nil{
		msgList = errCount.Error()
	} else {
		for fetchNotifications.Next(context.TODO()){
			err := fetchNotifications.Decode(&notifications)

			if err != nil {
				statusList = false
				msgList = errCount.Error()
			} else {
				statusList = true
				msgList = "Success"
				notificationsList = notifications.NotificationsList
			}
		}
	}


	if statusCount && statusList{
		status = true
		msg = "Success"
	} else {
		msg = msgList + "." + msgCount
	}
	
	return status, msg, count, notificationsList
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

// Fetch my bookmarks and reactions from database
func FetchMyBookmarksAndReactions(userID primitive.ObjectID)(status bool, msg string, bookmarks []primitive.ObjectID, reactions []primitive.ObjectID){

	var bookmarksDecode common.FetchUserProjectBookmarkStruct
	var reactionsDecode common.FetchUserProjectReactionStruct

	reactionStatus := false
	bookmarkStatus := false
	reactionMsg := ""
	bookmarkMsg := ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()

	fetchUserProjectBookmarks := client.Database(dbName).Collection(common.CONST_MO_BOOKMARKS)
	resultBookmarks, errBookmarks := fetchUserProjectBookmarks.Find(context.TODO(), bson.M{"userid" : userID}, options.Find().SetProjection(bson.M{"projectids": 1}))

	if errBookmarks != nil {
		bookmarkMsg = errBookmarks.Error()
	} else {
		for resultBookmarks.Next(context.TODO()) {
			err := resultBookmarks.Decode(&bookmarksDecode)
			if err != nil {
				bookmarkMsg = err.Error()
			} else {
				bookmarkStatus = true
				bookmarks = bookmarksDecode.ProjectIds
			}
		}
	}

	fetchUserProjectReactions := client.Database(dbName).Collection(common.CONST_MO_USER_PROJECT_REACTIONS)
	resultReactions, errReactions := fetchUserProjectReactions.Find(context.TODO(), bson.M{"userid" : userID}, options.Find().SetProjection(bson.M{"projectids" : 1}))

	if errReactions != nil {
		reactionMsg = errReactions.Error()
	} else {
		for resultReactions.Next(context.TODO()) {
			err := resultReactions.Decode(&reactionsDecode)
			if err != nil {
				reactionMsg = err.Error()
			} else {
				reactionStatus = true
				reactions = reactionsDecode.ProjectIds
			}
		}
	}


	if bookmarkStatus && reactionStatus {
		status = true
		msg = "Success"
	} else {
		status = false
		msg = "Failed." + bookmarkMsg + " " + reactionMsg
	}

	return status, msg, bookmarks, reactions
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

// Calculate time elapsed since project creation in nearest unit
func TimeElapsed(inputTime string)(nearestTimeUnit string){
	
	newTime, _ := time.Parse("Mon Jan 2 15:04:05 MST 2006", inputTime)

	timeElapsed := int64(time.Since(newTime).Seconds())


	if timeElapsed < 60 {
		return fmt.Sprintf("%ds", timeElapsed)
	} else if timeElapsed < 3600 {
		return fmt.Sprintf("%dm", timeElapsed/60)
	} else if timeElapsed < 86400 {
		return fmt.Sprintf("%dh", timeElapsed/3600)
	} else if timeElapsed < 604800 {
		return fmt.Sprintf("%dd", timeElapsed/86400)
	} else if timeElapsed < 2592000 {
		return fmt.Sprintf("%dw", timeElapsed/604800)
	} else if timeElapsed < 31536000 {
		return fmt.Sprintf("%dmo", timeElapsed/2592000)
	} else {
		return fmt.Sprintf("%dy", timeElapsed/31536000)
	}
}