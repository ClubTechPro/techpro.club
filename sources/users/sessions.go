package users

import (
	"context"
	"net/http"

	"techpro.club/sources/common"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserSession struct{
	UserID string `json:"userId"`
	SessionID string `json:"sessionId"`
}

// Get status, user id from session cookie
func getUserID(sessionId string) (status bool, errMsg string, userID string) {

	// Fetch userId from database
	client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
			
	type userStruct struct{
		Userid string `json:"userid"`
	}
	var result userStruct
	savedUserSession := client.Database(dbName).Collection(common.CONST_MO_USER_SESSIONS)
	err := savedUserSession.FindOne(context.TODO(), bson.M{"sessionid": sessionId}, options.FindOne().SetProjection(bson.M{"userid": 1, "_id": 0})).Decode(&result)

	if err != nil {
		status = false
		errMsg = err.Error()
		userID = ""
	} else {
		status = true
		errMsg = ""
		userID = result.Userid
	}
	
	return status, errMsg, userID
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

// Save user session in database
func SaveUserSession(userId, sessionId string) (status bool, errMsg string) {

	// Insert into database
	result := UserSession{userId, sessionId}
	
	client := common.Mongoconnect()
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
func ValidateSession(w http.ResponseWriter, r *http.Request)(status bool, userID string){
	ok, sessionID := GetSession(w, r)
	status = false
	if ok {
		okUser, _, userId := getUserID(sessionID)
		userID = userId

		if okUser {
			status = true
		}
	}
	return status, userID
}