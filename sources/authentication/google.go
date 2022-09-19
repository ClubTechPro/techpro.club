package authentication

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"techpro.club/sources/common"
)

// var googleOauthConfig = &oauth2.Config{
// 	RedirectURL:  "http://localhost:8080/contributors/google/callback",
// 	ClientID:     "150555465051-oqk98kejjufjvfbjh55q4nsd2hjvi9q5.apps.googleusercontent.com",
// 	ClientSecret: "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",
// 	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
// 	Endpoint:     google.Endpoint,
// }

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func GoogleContributorLoginHandler(w http.ResponseWriter, r *http.Request) {

	// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  common.GetGoogleContributorRedirectURI(),
		ClientID:     common.GetGoogleClientID(),
		ClientSecret: common.GetGoogleClientSecret(),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
	AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	validate that it matches the the state query parameter on your redirect callback.
	*/

	
	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func GoogleContributorCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {

	// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  common.GetGoogleContributorRedirectURI(),
		ClientID:     common.GetGoogleClientID(),
		ClientSecret: common.GetGoogleClientSecret(),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}