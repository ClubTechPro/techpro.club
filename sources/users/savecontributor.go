package users

import (
	"context"
	"net/http"
	"time"

	"techpro.club/sources/common"
	"techpro.club/sources/mailers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStruct struct {
	Email 		string `json:"email"`
	Name 		string `json:"name"`
	Location 	string `json:"location"`
	ImageLink 	string `json:"imageLink"`
	RepoUrl 	string `json:"repoUrl"`
	Source 		string `json:"source"`
	CreatedDate string `json:"createdDate"`
}

// Save a user to database and send a welcome email for first time users
func SaveUser(w http.ResponseWriter, r *http.Request, email, name, location, imageLink, repoUrl, source string )(status bool, msg string, objectID interface{}){
	
	client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveUserCollection := client.Database(dbName).Collection(common.CONST_MO_USERS)
	countUsers, errSearch := saveUserCollection.CountDocuments(context.TODO(), bson.M{"email": email, "source": common.CONST_GITHUB})

	if( errSearch != nil){
		status = false
		msg = errSearch.Error()
		objectID = nil
	} else {

		if (countUsers == 0){
			
			timestamp := time.Now()
			user := UserStruct{email, name, location, imageLink, repoUrl, source, timestamp.Format("2006-01-02 15:04:05")}
			insert, err := saveUserCollection.InsertOne(context.TODO(), user)
			if err != nil{
				status = false
				msg = err.Error()
				objectID = nil
			} else {
				status = true
				msg = ""
				objectID = insert.InsertedID
					
				// Code is the session
				session := r.URL.Query().Get("code")
				SaveUserSession(objectID.(primitive.ObjectID).Hex(), session)

				mailers.RegistrationEmail(email, name)
			}
		}

		
	}
	
	return status, msg, objectID
}
