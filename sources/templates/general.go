package templates

import (
	"context"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
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