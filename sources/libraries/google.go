package libraries

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"techpro.club/sources/common"
	"techpro.club/sources/pages/contributors"
	"techpro.club/sources/users"
)



const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func GoogleLoggedinHandler(w http.ResponseWriter, r *http.Request, googleData []byte, accessToken, userType, session string) {
	if string(googleData) == "" {
		// Unauthorized users get an unauthorized message
		fmt.Println("UNAUTHORIZED!")
		return
	}
	

	// Prettifying the json
	var prettyJSON bytes.Buffer
	// json.indent is a library utility function to prettify JSON indentation
	parserr := json.Indent(&prettyJSON, []byte(googleData), "", "\t")
	if parserr != nil {
		fmt.Println("JSON parse error")
	}

	var jsonMap map[string]interface{}
	json.Unmarshal(googleData, &jsonMap)

	// fmt.Println(googleData)
	// fmt.Println(jsonMap)
	// fmt.Println(jsonMap["email"])

	login := fmt.Sprintf("%s", jsonMap["id"])
	email := fmt.Sprintf("%s", jsonMap["email"])
	name := fmt.Sprintf("%s", jsonMap["name"])
	location := ""
	imageLink := fmt.Sprintf("%s", jsonMap["picture"])
	repoUrl := ""

	ok, _ := users.CheckUserExists(email)

	if (ok && userType == common.CONST_USER_CONTRIBUTOR) {
		// Save user session and redirect to feeds page
		users.SaveUser(w, r, login,email, name, location, imageLink, repoUrl, common.CONST_GOOGLE, userType, session)
	}  else {

		// Callback pages for contributors and projects
		
		contributors.CallBackGoogle(w,r)
		
		status, msg, _ := users.SaveUser(w, r, login, email, name, location, imageLink, repoUrl, common.CONST_GOOGLE, userType, session)
		if(!status){
			fmt.Println(msg)
		} 
	}
}

func GoogleContributorLoginHandler(w http.ResponseWriter, r *http.Request) {

	// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  common.GetGoogleContributorRedirectURI(),
		ClientID:     common.GetGoogleClientID(),
		ClientSecret: common.GetGoogleClientSecret(),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/youtube.readonly"},
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

	data, accessToken, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	fmt.Fprintf(w, "UserInfo: %s\n", data)


	// Set session cookie
	session := common.GenerateRandomSession(16)
	ok, _ := users.GetSession(w, r)
	if !ok {
		
		users.SetSessionCookie(w,r,session)
	}
	GoogleLoggedinHandler(w, r, data, accessToken, common.CONST_USER_CONTRIBUTOR, session )
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

func getUserDataFromGoogle(code string) (googleData []byte, accessToken string, err error) {

	// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  common.GetGoogleContributorRedirectURI(),
		ClientID:     common.GetGoogleClientID(),
		ClientSecret: common.GetGoogleClientSecret(),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/youtube.readonly"},
		Endpoint:     google.Endpoint,
	}

	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, "", fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, "", fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, token.AccessToken, nil
}