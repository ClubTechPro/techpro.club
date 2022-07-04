package templates

import (
	"net/http"

	"techpro.club/sources/common"
)

// Fetch user name and image from saved browser cookies
func FetchUsernameImage(w http.ResponseWriter, r *http.Request) (status bool, msg, userName, image string) {
	status = false
	msg = ""
	
	// user name cookie
	userNameCookie, err := r.Cookie(common.CONST_USER_NAME)

	if err != nil {
		status = false
		msg = err.Error()
		userName = ""
	} else {
		status = true
		msg = "Success"
		userName = userNameCookie.Value
	}

	// user image cookie
	imageCookie, err := r.Cookie(common.CONST_USER_IMAGE)

	if err != nil {
		status = false
		msg = err.Error()
		image = ""
	} else {
		status = true
		msg = "Success"
		image = imageCookie.Value
	}

	return status, msg, userName, image
}