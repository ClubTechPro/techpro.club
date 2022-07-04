package templates

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test Contains
func TestContains(t *testing.T) {
	var testString []string = []string{"test", "test2", "test3"}
	var testSubstring string = "tset"

	var testResult bool = Contains(testString, testSubstring)

	if !testResult {
		t.Errorf("contains() failed")
	}
}

// Test SliceToCsv
func TestSliceToCsv(t *testing.T) {
	var testString []string = []string{"test", "test2", "test3"}

	var testResult string = SliceToCsv(testString)

	if testResult != "test,test2,test3" {
		t.Errorf("sliceToCsv() failed")
	}
}

// Test FetchProjectDetails
func TestFetchProjectDetails(t *testing.T) {
	var testProjectID string = "12345"
	var testUserID primitive.ObjectID = primitive.NewObjectID()

	status, msg, _ := FetchProjectDetails(testProjectID, testUserID)

	if !status {
		t.Errorf(msg)
	}
}