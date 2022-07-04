package users

import (
	"net/http"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test getUserIDFromSession
func TestGetUserIDFromSession(t *testing.T) {
	testSessionId := "12345"
	status, errMsg, _ := getUserIDFromSession(testSessionId)

	if !status {
		t.Error(errMsg)
	}

}

// Test SaveUserDbSession
func TestSaveUserDbSession(t *testing.T) {
	testUserId := primitive.NewObjectID()
	testSessionId := "12345"
	testEmail := ""

	status, msg := SaveUserDbSession(testUserId, testSessionId, testEmail)

	if !status {
		t.Error(msg)
	}
}


// Test ValidateDbSession
func TestValidateDbSession(t *testing.T) {
	var testW http.ResponseWriter = http.ResponseWriter(nil)
	var testR *http.Request = &http.Request{}
	status, _ := ValidateDbSession(testW, testR)

	if !status {
		t.Error("Session not valid")
	}
}

// Test deleteDbSession
func TestDeleteDbSession(t *testing.T) {
	var testW http.ResponseWriter = http.ResponseWriter(nil)
	var testR *http.Request = &http.Request{}
	testSessionId := "12345"

	status, msg := deleteDbSession(testW, testR, testSessionId)

	if !status {
		t.Error(msg)
	}
}

// Test CheckUserExists
func TestCheckUserExists(t *testing.T) {
	testEmail := ""

	status, msg := CheckUserExists(testEmail)

	if !status {
		t.Error(msg)
	}
}
