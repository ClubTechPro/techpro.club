package contributors

import (
	"net/http"
	"testing"
)

func TestPreferences(t *testing.T){
	req, _ := http.NewRequest("POST", "/my_url", reader) //BTW check for error
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}