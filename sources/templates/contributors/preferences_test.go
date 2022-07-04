package contributors

import (
	"net/http"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test FetchPreferences
func TestFetchPreferences(t *testing.T){
	var sampleUserId primitive.ObjectID = primitive.NewObjectID()
	
	status, msg, _ := fetchPreferences(sampleUserId)

	if !status {
		t.Errorf(msg)
	}
}

// Test SavePreferences
func TestSavePreferences(t *testing.T){
	var testUserId primitive.ObjectID = primitive.NewObjectID()
	var testW http.ResponseWriter = http.ResponseWriter(nil)
	var testR *http.Request = new(http.Request)

	status, msg := savePreferences(testW, testR, testUserId)
	
	if !status {
		t.Errorf(msg)
	}
}