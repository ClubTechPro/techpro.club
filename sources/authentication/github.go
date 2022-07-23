package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/templates/contributors"
	"techpro.club/sources/templates/projects"
	"techpro.club/sources/users"
)

type GithubEmailStruct struct{
	Email string `json:"email"`
	Primary bool `json:"primary"`
	Verified bool `json:"verified"`
	Visibility string `json:"visibility"`
}

func GithubLoggedinHandler(w http.ResponseWriter, r *http.Request, githubData, accessToken, userType string) {
	if githubData == "" {
		// Unauthorized users get an unauthorized message
		fmt.Println("UNAUTHORIZED!")
		return
	}
	

	// Prettifying the json
	var prettyJSON bytes.Buffer
	// json.indent is a library utility function to prettify JSON indentation
	parserr := json.Indent(&prettyJSON, []byte(githubData), "", "\t")
	if parserr != nil {
		fmt.Println("JSON parse error")
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(githubData), &jsonMap)

	login := fmt.Sprintf("%s", jsonMap["login"])
	email := fmt.Sprintf("%s", jsonMap["email"])
	name := fmt.Sprintf("%s", jsonMap["name"])
	location := fmt.Sprintf("%s", jsonMap["location"])
	imageLink := fmt.Sprintf("%s", jsonMap["avatar_url"])
	repoUrl := fmt.Sprintf("%s", jsonMap["html_url"])

	// Github sends "%!s(<nil>)", if nil found
	fmt.Println("ACCESS", accessToken)
	if email == "%!s(<nil>)"{
		
		email = GetUserEmail(accessToken)
	} 

	ok, _ := users.CheckUserExists(email)

	if (ok && userType == common.CONST_USER_CONTRIBUTOR) {
		// Save user session and redirect to feeds page
		users.SaveUser(w, r, login,email, name, location, imageLink, repoUrl, common.CONST_GITHUB, userType)
	} else if (ok && userType == common.CONST_USER_PROJECT) {
		// Save user session and redirect to projects list page
		users.SaveUser(w, r, login, email, name, location, imageLink, repoUrl, common.CONST_GITHUB, userType)
	} else {

		// Callback pages for contributors and projects
		
		if userType == common.CONST_USER_CONTRIBUTOR{
			contributors.CallBack(w,r)
		} else {
			projects.CallBack(w,r)
		}
		
		status, msg, _ := users.SaveUser(w, r, login, email, name, location, imageLink, repoUrl, common.CONST_GITHUB, userType)
		if(!status){
			fmt.Println(msg)
		} 
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

func GithubContributorCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Set session cookie
	ok, _ := users.GetSession(w, r)
	if !ok {
		users.SetSessionCookie(w,r,code)
	}
	
	githubAccessToken := GetGithubAccessToken(code)

	githubData := GetGithubData(githubAccessToken)

	GithubLoggedinHandler(w, r, githubData, githubAccessToken, common.CONST_USER_CONTRIBUTOR)
}

func GithubProjectCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Set session cookie
	ok, _ := users.GetSession(w, r)
	if !ok {
		users.SetSessionCookie(w,r,code)
	}

	githubAccessToken := GetGithubAccessToken(code)

	githubData := GetGithubData(githubAccessToken)

	GithubLoggedinHandler(w, r, githubData, githubAccessToken, common.CONST_USER_PROJECT)
}

func GetGithubData(accessToken string) (responseBody string) {
	req, reqerr := http.NewRequest("GET", "https://api.github.com/user", nil)
	if reqerr != nil {
		fmt.Println("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		fmt.Println("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	return string(respbody)
}

func GetUserEmail(accessToken string)(primaryEmail string){
	req, reqerr := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if reqerr != nil {
		fmt.Println("API Request creation failed")
	}

	fmt.Println("accessToken", accessToken)
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		fmt.Println("Request failed")
	} 

	respbody, _ := ioutil.ReadAll(resp.Body)

	var emailStructSlice []GithubEmailStruct
    err := json.Unmarshal(respbody, &emailStructSlice)
    if err != nil {
        panic(err)
    } else {
		for _, emailStruct := range emailStructSlice {

			if (emailStruct.Primary){
				primaryEmail = emailStruct.Email
			}
		}
	}

	return primaryEmail
}

func GetGithubAccessToken(code string) (accessToken string) {

	clientID := common.GetGithubClientID()
	clientSecret := common.GetGithubClientSecret()

	requestBodyMap := map[string]string{"client_id": clientID, "client_secret": clientSecret, "code": code}
	requestJSON, _ := json.Marshal(requestBodyMap)

	req, reqerr := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON))
	if reqerr != nil {
		fmt.Println("Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		fmt.Println("Request failed")
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