package users

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"techpro.club/sources/common"
	"techpro.club/sources/mailers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStruct struct {
	Login       string `json:"login"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	ImageLink   string `json:"imageLink"`
	RepoUrl     string `json:"repoUrl"`
	Source      string `json:"source"`
	CreatedDate string `json:"createdDate"`
}

// Save a user to database and send a welcome email for first time users
func SaveUser(w http.ResponseWriter, r *http.Request, login, email, name, location, imageLink, repoUrl, source, userType, session string) (status bool, msg string, userIdObject interface{}) {

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	userCollection := client.Database(dbName).Collection(common.CONST_MO_USERS)
	countUsers, errSearch := userCollection.CountDocuments(context.TODO(), bson.M{"email": email})

	if errSearch != nil {
		status = false
		msg = errSearch.Error()
		userIdObject = nil
	} else {

		if countUsers == 0 {

			timestamp := time.Now()
			user := UserStruct{login, email, name, location, imageLink, repoUrl, source, timestamp.Format("2006-01-02 15:04:05")}
			insert, err := userCollection.InsertOne(context.TODO(), user)
			if err != nil {
				status = false
				msg = err.Error()
				userIdObject = nil
			} else {
				status = true
				msg = ""
				userIdObject = insert.InsertedID

				// Code is the session
				// session := r.URL.Query().Get("code")
				SaveUserDbSession(userIdObject.(primitive.ObjectID), session, email)
				SetUserCookie(w, r, name)
				SetUserImageCookie(w, r, imageLink)

				mailers.RegistrationEmail(email, name, userType)
			}
		} else {

			type useIdStruct struct {
				ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
			}
			var result useIdStruct
			err := userCollection.FindOne(context.TODO(), bson.M{"email": email}, options.FindOne()).Decode(&result)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				// session := r.URL.Query().Get("code")
				SaveUserDbSession(result.ID, session, email)
				SetUserCookie(w, r, name)
				SetUserImageCookie(w, r, imageLink)
			}

			// Redirect to respective home pages of user
			if userType == common.CONST_USER_CONTRIBUTOR {
				http.Redirect(w, r, "/contributors/feeds", http.StatusPermanentRedirect)
			} else {
				http.Redirect(w, r, "/projects/list", http.StatusPermanentRedirect)
			}
		}

	}

	return status, msg, userIdObject
}
