package projects

import (
	"net/http"
	"testing"

	"techpro.club/sources/common"
)

// Test UpdateProject
func TestUpdateProject(t *testing.T){
	var testProjectID string = "62bd7328bf850f09cb4d5a3a"
	var newProjectStruct common.SaveProjectStruct = common.SaveProjectStruct{}
	var testW http.ResponseWriter = http.ResponseWriter(nil)
	var testR *http.Request = new(http.Request)

	status, msg := updateProject(testW, testR, testProjectID, newProjectStruct)

	if !status {
		t.Errorf(msg)
	}
}