package contributors

import (
	"testing"
)

func TestFilterActiveProjects(t *testing.T){
	
	// TEST CONDITIONS
	// This has to come from the actual frontend
	pageid := int64(0)
	tags := []string{"cpp"}
	keyword := "Go"

	status, msg, _ := filterActiveProjects(pageid, tags, keyword)

	if !status {
		t.Errorf(msg)
	}
}