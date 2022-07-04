package templates

import (
	"net/http"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test fetchNotificationList
func TestFetchNotificationList(t *testing.T){
	var testUserID primitive.ObjectID = primitive.NewObjectID()
	var testW http.ResponseWriter = http.ResponseWriter(nil)
	var testR *http.Request = new(http.Request)
	
	status, msg, _ := fetchNotificationList(testW, testR, testUserID)
	
	if !status {
		t.Errorf(msg)
	}
}