package contributors

import (
	"net/http"
	"testing"
)

func TestCallBack(t *testing.T){
	
	// Check for get request
	_, errGet := http.NewRequest(http.MethodGet, "/contributor/github/callback", nil) 
	if errGet != nil {
        t.Errorf(errGet.Error())
    }
}