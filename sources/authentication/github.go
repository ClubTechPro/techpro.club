package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/templates/contributors"
	"techpro.club/sources/users"
)


func GithubLoggedinHandler(w http.ResponseWriter, r *http.Request, githubData string) {
	if githubData == "" {
		// Unauthorized users get an unauthorized message
		fmt.Println("UNAUTHORIZED!")
		return
	}

	w.Header().Set("Content-type", "application/json")

	// Prettifying the json
	var prettyJSON bytes.Buffer
	// json.indent is a library utility function to prettify JSON indentation
	parserr := json.Indent(&prettyJSON, []byte(githubData), "", "\t")
	if parserr != nil {
		log.Panic("JSON parse error")
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(githubData), &jsonMap)

	email := fmt.Sprintf("%s", jsonMap["email"])
	name := fmt.Sprintf("%s", jsonMap["name"])
	location := fmt.Sprintf("%s", jsonMap["location"])
	imageLink := fmt.Sprintf("%s", jsonMap["avatar_url"])
	repoUrl := fmt.Sprintf("%s", jsonMap["html_url"])

	
	// Return the prettified JSON as a string
	// fmt.Fprintf(w, string(prettyJSON.Bytes()))
	status, msg, _ := users.SaveUser(w, r, email, name, location, imageLink, repoUrl, common.CONST_GITHUB)
	
	if(!status){
		fmt.Println(msg)
	} 
}

func GithubContributorLoginHandler(w http.ResponseWriter, r *http.Request) {

	githubClientID := common.GetGithubClientID()
	githubContributorRedirectUri := common.GetGithubContributorRedirectURI()

	redirectURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", githubClientID, githubContributorRedirectUri)

	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
}

func GithubProjectLoginHandler(w http.ResponseWriter, r *http.Request) {
	githubClientID := common.GetGithubClientID()
	githubProjectRedirectUri := common.GetGithubProjectRedirectURI()

	redirectURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", githubClientID, githubProjectRedirectUri)

	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
}

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Set session cookie
	ok, _ := users.GetSession(w, r)
	if !ok {
		users.SetSessionCookie(w,r,code)
	}
	
	contributors.CallBack(w,r)
	githubAccessToken := GetGithubAccessToken(code)

	githubData := GetGithubData(githubAccessToken)

	GithubLoggedinHandler(w, r, githubData)
}

func GetGithubData(accessToken string) (responseBody string) {
	req, reqerr := http.NewRequest("GET", "https://api.github.com/user", nil)
	if reqerr != nil {
		log.Panic("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	return string(respbody)
}

func GetGithubAccessToken(code string) (accessToken string) {

	clientID := common.GetGithubClientID()
	clientSecret := common.GetGithubClientSecret()

	requestBodyMap := map[string]string{"client_id": clientID, "client_secret": clientSecret, "code": code}
	requestJSON, _ := json.Marshal(requestBodyMap)

	req, reqerr := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON))
	if reqerr != nil {
		log.Panic("Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	// Represents the response received from Github
	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	var ghresp githubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	return ghresp.AccessToken
}