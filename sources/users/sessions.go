package users

import (
	"context"
	"fmt"
	"net/http"

	"techpro.club/sources/common"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserSession struct{
	UserID primitive.ObjectID `json:"userId"`
	SessionID string `json:"sessionId"`
}


// Get status, user id from session cookie
func getUserIDFromSession(sessionId string) (status bool, errMsg string, userID primitive.ObjectID) {

	// Fetch userId from database
	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
			
	type userStruct struct{
		Userid primitive.ObjectID `json:"userid"`
	}
	var result userStruct
	savedUserSession := client.Database(dbName).Collection(common.CONST_MO_USER_SESSIONS)
	err := savedUserSession.FindOne(context.TODO(), bson.M{"sessionid": sessionId}, options.FindOne().SetProjection(bson.M{"userid": 1, "_id": 0})).Decode(&result)

	if err != nil {
		status = false
		errMsg = err.Error()
		userID = primitive.NilObjectID
	} else {
		status = true
		errMsg = ""
		userID = result.Userid
	}
	
	return status, errMsg, userID
}

// Set session cookie for user
func SetSessionCookie(w http.ResponseWriter, r *http.Request, session string) (sessionID string) {
	
	// session cookie
	sessionCookie := &http.Cookie{
		Name : common.CONST_SESSION_NAME,
		Value: session,
		Path : "/",
	}
	http.SetCookie(w, sessionCookie)
	return sessionID
}


// Get user session from cookie
func GetSession(w http.ResponseWriter, r *http.Request) (status bool, sessionID string) {

	sessionCookie, err := r.Cookie(common.CONST_SESSION_NAME)

	if err != nil {
		status = false
		sessionID = ""
	} else {
		status = true
		sessionID = sessionCookie.Value
	}
	return status, sessionID
}

// Delete session cookie for user
func DeleteSessionCookie(w http.ResponseWriter, r *http.Request) {
	
	// Get current session
	ok, sessionID := GetSession(w, r)

	if ok {
		status, errMsg := deleteDbSession(w, r, sessionID)

		if !status {
			fmt.Println(errMsg)
		}
	}
	// delete session cookie
	sessionCookie := &http.Cookie{
		Name : common.CONST_SESSION_NAME,
		Value: "",
		Path : "/",
		MaxAge:-1,
	}
	http.SetCookie(w, sessionCookie)
}

/* User name session methods*/
// Set user name cookie
func SetUserCookie(w http.ResponseWriter, r *http.Request, userName string) {
	// user cookie
	userCookie := &http.Cookie{
		Name : common.CONST_USER_NAME,
		Value: userName,
		Path : "/",
	}
	http.SetCookie(w, userCookie)
}

// Get user name from cookie
func GetUserName(w http.ResponseWriter, r *http.Request) (status bool, userName string) {

	userNameCookie, err := r.Cookie(common.CONST_USER_NAME)

	if err != nil {
		status = false
		userName = ""
	} else {
		status = true
		userName = userNameCookie.Value
	}
	return status, userName
}

// Delete user name cookie
func DeleteUserCookie(w http.ResponseWriter, r *http.Request) {
	// user cookie
	userCookie := &http.Cookie{
		Name : common.CONST_USER_NAME,
		Value: "",
		Path : "/",
		MaxAge:-1,
	}
	http.SetCookie(w, userCookie)
}

/* User image session methods*/
// Set image cookie
func SetUserImageCookie(w http.ResponseWriter, r *http.Request, image string) {
	// user cookie
	userCookie := &http.Cookie{
		Name : common.CONST_USER_IMAGE,
		Value: image,
		Path : "/",
	}
	http.SetCookie(w, userCookie)
}

// Get image cookie
func GetUserImageCookie(w http.ResponseWriter, r *http.Request) (status bool, userImage string) {
	userImageCookie, err := r.Cookie(common.CONST_USER_IMAGE)

	if err != nil {
		status = false
		userImage = ""
	} else {
		status = true
		userImage = userImageCookie.Value
	}
	return status, userImage
}


// Save user session in database
func SaveUserDbSession(userId primitive.ObjectID, sessionId, email string) (status bool, errMsg string) {

	// Insert into database
	result := UserSession{userId, sessionId}
	
	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveUserSession := client.Database(dbName).Collection(common.CONST_MO_USER_SESSIONS)

	_, err := saveUserSession.InsertOne(context.TODO(), result)

	if err != nil {
		status = false
		errMsg = err.Error()

	} else {
		status = true
		errMsg = ""
	}

	return status, errMsg
}

// Checks if session is valid, else return false
func ValidateDbSession(w http.ResponseWriter, r *http.Request)(status bool, userID primitive.ObjectID){
	ok, sessionID := GetSession(w, r)
	status = false

	if ok {
		okUser, _, userId := getUserIDFromSession(sessionID)
		userID = userId

		if okUser {
			status = true
		}
	} 
	return status, userID
}

// Delete the session from the databse
func deleteDbSession(w http.ResponseWriter, r *http.Request, sessionID string)(status bool, errMsg string) {

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	_, userID := ValidateDbSession(w, r)

	dbName := common.GetMoDb()
	saveUserSession := client.Database(dbName).Collection(common.CONST_MO_USER_SESSIONS)

	_, err := saveUserSession.DeleteMany(context.TODO(), bson.M{"userid": userID})

	if err != nil {
		status = false
		errMsg = err.Error()
	} else {
		status = true
		errMsg = ""
	}
	
	return status, errMsg
}

// Check if user exists
func CheckUserExists(email string)(status bool, msg string){
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	countUserExists := client.Database(dbName).Collection(common.CONST_MO_USERS)

	countUsers, err := countUserExists.CountDocuments(context.TODO(), bson.M{"email": email})

	if err != nil {
		msg = err.Error()
	} else {
		if countUsers > 0 {
			status = true
		} else {
			msg = "User does not exist"
		}
	}
	

	return status, msg
}
