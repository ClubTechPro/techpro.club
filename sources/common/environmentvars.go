package common

import (
	"log"
	"os"
)

// Get base url from environment variable
func GetBaseurl() (baseUrl string) {

	baseUrl, exists := os.LookupEnv("BASE_URL")
	if !exists {
		log.Fatal("BASE_URL not defined in .env file")
	}

	return baseUrl
}

// Get mongodb host from environment variable
func GetMoHost() (moHost string) {

	moHost, exists := os.LookupEnv("MO_HOST")
	if !exists {
		log.Fatal("MO_HOST not defined in .env file")
	}

	return moHost
}

// Get mongodb port from environment variable
func GetMoPort() (moPort string) {

	moPort, exists := os.LookupEnv("MO_PORT")
	if !exists {
		log.Fatal("MO_PORT not defined in .env file")
	}

	return moPort
}

// Get mongodb db name from environment variable
func GetMoDb() (moDb string) {

	moDb, exists := os.LookupEnv("MO_DATABASE")
	if !exists {
		log.Fatal("MO_DATABASE not defined in .env file")
	}

	return moDb
}

// Get mongodb user from environment variable
func GetMoUser() (moUser string) {

	moUser, exists := os.LookupEnv("MO_USER")
	if !exists {
		log.Fatal("MO_USER not defined in .env file")
	}

	return moUser
}

// Get mongodb password from environment variable
func GetMoPass() (moPass string) {

	moPass, exists := os.LookupEnv("MO_PASS")
	if !exists {
		log.Fatal("MO_PASS not defined in .env file")
	}

	return moPass
}

// Get mongodb auth db MO_AUTH_DB from environment variable
func GetMoAuthDb() (moAuthDb string) {

	moAuthDb, exists := os.LookupEnv("MO_AUTH_DB")
	if !exists {
		log.Fatal("MO_AUTH_DB not defined in .env file")
	}

	return moAuthDb
}

// Get mongodb auth method MO_AUTH_METHOD from environment variable
func GetMoAuthMethod() (moAuthMethod string) {

	moAuthMethod, exists := os.LookupEnv("MO_AUTH_METHOD")
	if !exists {
		log.Fatal("MO_AUTH_METHOD not defined in .env file")
	}

	return moAuthMethod
}

// Get Google client id from environment variable
func GetGoogleClientID() (googleClientID string) {

	googleClientID, exists := os.LookupEnv("GO_CLIENT_ID")
	if !exists {
		log.Fatal("Google Client ID not defined in .env file")
	}

	return googleClientID
}

// Get Google client secret from environment variable
func GetGoogleClientSecret() (googleClientSecret string) {

	googleClientSecret, exists := os.LookupEnv("GO_CLIENT_SECRET")
	if !exists {
		log.Fatal("Google Client Secret not defined in .env file")
	}

	return googleClientSecret
}

// Get Google contributor redirect uri from environment variable
func GetGoogleContributorRedirectURI() (googleContributorRedirectURI string) {

	googleContributorRedirectURI, exists := os.LookupEnv("GO_CONTRIBUTOR_REDIRECT_URI")
	if !exists {
		log.Fatal("Google Contributor Redirect URI not defined in .env file")
	}

	return googleContributorRedirectURI
}

// Get Github client id from environment variable
func GetGithubClientID() (githubClientID string) {

	githubClientID, exists := os.LookupEnv("GB_CLIENT_ID")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}

	return githubClientID
}

// Get Github client secret from environment variable
func GetGithubClientSecret() (githubClientSecret string) {

	githubClientSecret, exists := os.LookupEnv("GB_CLIENT_SECRET")
	if !exists {
		log.Fatal("Github Client Secret not defined in .env file")
	}

	return githubClientSecret
}

// Get Github contributor redirect uri from environment variable
func GetGithubContributorRedirectURI() (githubContributorRedirectURI string) {

	githubContributorRedirectURI, exists := os.LookupEnv("GB_CONTRIBUTOR_REDIRECT_URI")
	if !exists {
		log.Fatal("Github Contributor Redirect URI not defined in .env file")
	}

	return githubContributorRedirectURI
}

// Get Github project redirect uri from environment variable
func GetGithubProjectRedirectURI() ( githubProjectRedirectURI string) {

	githubProjectRedirectURI, exists := os.LookupEnv("GB_PROJECT_REDIRECT_URI")
	if !exists {
		log.Fatal("Github Project Redirect URI not defined in .env file")
	}

	return githubProjectRedirectURI
}


// Get AWS SES region from environment variable
func GetSesRegion() (sesRegion string) {

	sesRegion, exists := os.LookupEnv("SES_REGION")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}

	return sesRegion
}

// Get AWS SES AccessID from environment variable
func GetSesAccessID() (sesAccessID string) {

	sesAccessID, exists := os.LookupEnv("SES_ACCESS_ID")
	if !exists {
		log.Fatal("AWS SES AccessID not defined in .env file")
	}

	return sesAccessID
}

// Get AWS SES SecretKey from environment variable
func GetSesSecretKey() ( sesSecretKey string) {

	sesSecretKey, exists := os.LookupEnv("SES_ACCESS_SECRET")
	if !exists {
		log.Fatal("AWS SES SecretKey  not defined in .env file")
	}

	return sesSecretKey
}

// Get AWS SES SecretKey from environment variable
func GetSesSender() ( sesSender string) {

	sesSender, exists := os.LookupEnv("SES_SENDER")
	if !exists {
		log.Fatal("AWS SES Sender  not defined in .env file")
	}

	return sesSender
}