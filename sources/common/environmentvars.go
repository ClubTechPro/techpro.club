package common

import (
	"log"
	"os"
)

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

// Get Github client id from environment variable
func GetGithubClientID() (githubClientID string) {

	githubClientID, exists := os.LookupEnv("GITHUB_CLIENT_ID")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}

	return githubClientID
}

// Get Github client secret from environment variable
func GetGithubClientSecret() (githubClientSecret string) {

	githubClientSecret, exists := os.LookupEnv("GITHUB_CLIENT_SECRET")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}

	return githubClientSecret
}

// Get Github contributor redirect uri from environment variable
func GetGithubContributorRedirectURI() (githubContributorRedirectURI string) {

	githubContributorRedirectURI, exists := os.LookupEnv("GITHUB_CONTRIBUTOR_REDIRECT_URI")
	if !exists {
		log.Fatal("Github Contributor Redirect URI not defined in .env file")
	}

	return githubContributorRedirectURI
}

// Get Github project redirect uri from environment variable
func GetGithubProjectRedirectURI() ( githubProjectRedirectURI string) {

	githubProjectRedirectURI, exists := os.LookupEnv("GITHUB_PROJECT_REDIRECT_URI")
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