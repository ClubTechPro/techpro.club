package main

import (
	"log"
	"net/http"

	"techpro.club/sources/authentication"
	"techpro.club/sources/common"
	"techpro.club/sources/notifications"
	"techpro.club/sources/templates"
	"techpro.club/sources/templates/contributors"
	"techpro.club/sources/templates/projects"

	"github.com/joho/godotenv"
)


func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
    fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// Templates
	http.HandleFunc("/", templates.IndexHandler)
	http.HandleFunc("/contactus", templates.ContactUs)
	http.HandleFunc("/careers", templates.Careers)
	http.HandleFunc("/company", templates.Company)
	http.HandleFunc("/brand", templates.Brand)
	http.HandleFunc("/campus", templates.Campus)
	http.HandleFunc("/campusonboard", templates.CampusOnboard)

	// Users
	http.HandleFunc("/users/settings", templates.UserSettings)
	http.HandleFunc("/users/notifications", notifications.Fetch)

	// Templates/Contributors
	http.HandleFunc("/contributors/feeds", contributors.Feeds)
	http.HandleFunc("/contributors/preferences", contributors.Preferences)
	http.HandleFunc("/contributors/thankyou", contributors.PreferencesSaved)
	

	// Templates/Contributors
	http.HandleFunc("/projects/create", projects.ProjectCreate)
	http.HandleFunc("/projects/list", projects.ProjectList)
	http.HandleFunc("/projects/view", projects.ProjectPreview)
	http.HandleFunc("/projects/edit", projects.ProjectEdit)
	http.HandleFunc("/projects/thankyou", projects.ProjectSaved)


	// Authentication
	// Github
	http.HandleFunc("/contributors/github/login/", authentication.GithubContributorLoginHandler)
	http.HandleFunc("/contributors/github/callback", authentication.GithubContributorCallbackHandler)

	http.HandleFunc("/projects/github/login/", authentication.GithubProjectLoginHandler)
	http.HandleFunc("/projects/github/callback", authentication.GithubProjectCallbackHandler)

	// Func to receive data after login
	http.HandleFunc("/github/loggedin", func(w http.ResponseWriter, r *http.Request) {
		authentication.GithubLoggedinHandler(w, r, "", "", "")
	})

	http.HandleFunc("/logout", templates.Logout)

	// Start the web server
    http.ListenAndServe(common.CONST_APP_PORT, nil)
}
