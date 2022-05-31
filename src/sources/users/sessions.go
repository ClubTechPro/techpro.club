package users

import (
	"config"
	"context"
	"fmt"
	"net/http"
	"sources/common"

	"github.com/satori/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type UserSession struct{
	UserID string `json:"userId"`
	SessionID string `json:"sessionId"`
}

// Get status, user id from session cookie
func getUserID(sessionId string) (status bool, errMsg string, userID string) {

	// Fetch userId from database
	client := config.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
			
	var result bson.M
	savedUserSession := client.Database(dbName).Collection(common.CONST_MO_USER_SESSIONS)
	err := savedUserSession.FindOne(context.TODO(), bson.M{"sessionId": sessionId}).Decode(&result)

	if err != nil {
		status = false
		errMsg = err.Error()
		userID = ""
		fmt.Println("Error Session find", err)
	} else {
		status = true
		errMsg = ""
		userID = ""
	}
	fmt.Println(result)

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
func SetSessionCookie(w http.ResponseWriter, r *http.Request) (sessionID string) {
	
	// session cookie
	uid := uuid.Must(uuid.NewV4())
	sessionID = uid.String()
	
	sessionCookie := &http.Cookie{
		Name : common.CONST_SESSION_NAME,
		Value: sessionID,
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
	}
	http.SetCookie(w, userCookie)
}

// Validate cookie for user
// If not exists or invalid, create a new session cookie
func ValidateCookie(w http.ResponseWriter, r *http.Request) {

	_, errSessionCookie := r.Cookie(common.CONST_SESSION_NAME)

	if errSessionCookie != nil {
		// Set a new cookie and redirect
		SetSessionCookie(w, r)
	} 
}

// Save user session in database
func SaveUserSession(userId, sessionId string) (status bool, errMsg string) {

	// Insert into database
	result := UserSession{userId, userId}
	
	client := config.Mongoconnect()
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