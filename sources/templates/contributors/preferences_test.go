package contributors

import (
	"net/http"
	"strings"
	"testing"
)

func TestPreferences(t *testing.T){

	// Check for get request
	_, errGet := http.NewRequest(http.MethodGet, "/contributor/preferences", nil) 
	if errGet != nil {
        t.Errorf(errGet.Error())
    }

	// Check for post request
	req, errPost := http.NewRequest(http.MethodPost, "/contributor/preferences",
		strings.NewReader("notificationFrequency=a&contributorCount=b"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if errPost != nil {
        t.Errorf(errPost.Error())
    }
}