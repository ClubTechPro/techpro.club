package templates

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test fetchUserProfile
func TestFetchUserProfile(t *testing.T){
	var testUserID primitive.ObjectID = primitive.NewObjectID()

	status, msg, _ := fetchUserProfile(testUserID)
	if !status {
		t.Errorf(msg)
	}
}


// Test UpdateUserProfile
func TestUpdateUserProfile(t *testing.T){
	var testUserID primitive.ObjectID = primitive.NewObjectID()
	var testName, testRepoLink string

	status, msg := UpdateUserProfile(testUserID, testName, testRepoLink)
	if !status {
		t.Errorf(msg)
	}
}

// Test fetchSocials
func TestFetchSocials(t *testing.T){
	var testUserID primitive.ObjectID = primitive.NewObjectID()

	status, msg, _ := fetchSocials(testUserID)
	if !status {
		t.Errorf(msg)
	}
}


// Test updateSocials
func TestUpdateSocials(t *testing.T){
	var testUserID primitive.ObjectID  = primitive.NewObjectID()
	var testFacebook, testLinkedin, testTwitter, testStackoverflow string

	status, msg := updateSocials(testUserID, testFacebook, testLinkedin, testTwitter, testStackoverflow)

	if !status {
		t.Errorf(msg)
	}
}