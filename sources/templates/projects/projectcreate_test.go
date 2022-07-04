package projects

import (
	"net/http"
	"testing"

	"techpro.club/sources/common"
)

// Test SaveProject
func TestSaveProject(t *testing.T){
	var testProjectStruct common.SaveProjectStruct = common.SaveProjectStruct{}
	var testW http.ResponseWriter = http.ResponseWriter(nil)
	var testR *http.Request = new(http.Request)

	status, msg := saveProject(testW, testR, testProjectStruct)
	
	if !status {
		t.Errorf(msg)
	}
}