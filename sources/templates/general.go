package templates

import (
	"net/http"

	"techpro.club/sources/common"
)

// Fetch user name and image from saved browser cookies
func FetchUsernameImage(w http.ResponseWriter, r *http.Request) (status bool, userName, image string) {
	// user cookie
	userNameCookie, err := r.Cookie(common.CONST_USER_NAME)

	if err != nil {
		status = false
		userName = ""
	} else {
		status = true
		userName = userNameCookie.Value
	}

	// user cookie
	imageCookie, err := r.Cookie(common.CONST_USER_IMAGE)

	if err != nil {
		status = false
		image = ""
	} else {
		status = true
		image = imageCookie.Value
	}

	return status, userName, image
}