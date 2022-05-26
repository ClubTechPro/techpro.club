package users

import (
	"config"
	"context"
	"sources/common"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func SaveUser(email, name, location, imageLink, repoUrl, source string )(status bool, msg string, objectID interface{}){

	client := config.Mongoconnect()
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
			user := UserStruct{"ded", name, location, imageLink, repoUrl, source, timestamp.Format("2006-01-02 15:04:05")}
			insert, err := saveUserCollection.InsertOne(context.TODO(), user)
			if err != nil{
				status = false
				msg = err.Error()
				objectID = nil
			} else {
				status = true
				msg = ""
				objectID = insert.InsertedID
			}
		}
	}
	
	return status, msg, objectID
}
