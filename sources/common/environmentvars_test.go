package common

import (
	"testing"
)

func TestGetGithubClientID(t *testing.T){
	got := GetGithubClientID()

    if got != "" {
		t.Errorf("Github client missing")

	} 
}

func TestGetGithubClientSecret(t *testing.T){
	got := GetGithubClientSecret()

    if got != "" {
		t.Errorf("Github secret missing")

	} 
}

func TestGetGithubContributorRedirectURI(t *testing.T){
	got := GetGithubContributorRedirectURI()

    if got != "" {
		t.Errorf("Contributor Redirect URI missing")

	} 
}

func TestGetGithubProjectRedirectURI(t *testing.T){
	got := GetGithubProjectRedirectURI()

    if got != "" {
		t.Errorf("Project Redirect URI missing")

	} 
}

// Get AWS SES region from environment variable
func TestGetSesRegion(t *testing.T) {

	got := GetSesRegion()

    if got != "" {
		t.Errorf("SES region missing")

	} 
}

// Get AWS SES AccessID from environment variable
func TestGetSesAccessID(t *testing.T)  {

	got := GetSesAccessID()

    if got != "" {
		t.Errorf("SES Access ID missing")

	} 
}

// Get AWS SES SecretKey from environment variable
func TestGetSesSecretKey(t *testing.T){

	got := GetSesSecretKey()

    if got != "" {
		t.Errorf("SES secret key missing")

	} 
}

// Get AWS SES SecretKey from environment variable
func TestGetSesSender(t *testing.T){

	got := GetSesSender()

    if got != "" {
		t.Errorf("SES sender missing")

	} 
}