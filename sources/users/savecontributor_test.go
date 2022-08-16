package users

import (
	"net/http"
	"testing"
)

// Test SaveUser
func TestSaveUser(t *testing.T) {
	var testEmail, testName, testLogin, testLocation, testImageLink, testRepoUrl, testSource, testUserType string
	var testW http.ResponseWriter = http.ResponseWriter(nil)
	var testR *http.Request = new(http.Request)

	status, msg, _ := SaveUser(testW, testR, testEmail, testName, testLocation, testImageLink, testRepoUrl, testSource, testUserType, testLogin)

	if !status {
		t.Errorf(msg)
	}
}
