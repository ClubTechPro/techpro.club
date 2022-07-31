package projects

import (
	"net/http"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestListProjects(t *testing.T){
	var testUserID primitive.ObjectID = primitive.NewObjectID()
	var testW http.ResponseWriter = http.ResponseWriter(nil)
	var testR *http.Request = new(http.Request)

	status, msg, _ := listProjects(testW, testR, testUserID)
	
	if !status {
		t.Errorf(msg)
	}
}