package templates

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test Contains
func TestContains(t *testing.T) {
	var testString []string = []string{"test", "test2", "test3"}
	var testSubstring string = "tset"
	var testSubstring2 string = "test"

	// Check pass
	status := Contains(testString, testSubstring)

	if status {
		t.Errorf("contains() failed")
	}

	// Check fail
	status = Contains(testString, testSubstring2)

	if !status {
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
	var testProjectID string = "62bd7328bf850f09cb4d5a3a"
	testUserID := primitive.NewObjectID()

	status, msg, _ := FetchProjectDetails(testProjectID, testUserID)

	if !status {
		t.Errorf(msg)
	}
}

// Test NotificationsCount
func TestNotificationsCount(t *testing.T) {
	var testUserID primitive.ObjectID = primitive.NewObjectID()

	status, msg, _ := NotificationsCount(testUserID)

	if !status {
		t.Errorf(msg)
	}
}