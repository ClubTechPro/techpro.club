package main

import (
	"log"
	"net/http"

	"techpro.club/sources/authentication"
	"techpro.club/sources/common"
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
	http.HandleFunc("/projects", templates.ProjectIndexHandler)

	// Templates/Contributors
	http.HandleFunc("/contributor/preferences", contributors.Preferences)
	http.HandleFunc("/contributor/thankyou", contributors.PreferencesSaved)
	

	// Templates/Contributors
	http.HandleFunc("/projects/create", projects.ProjectCreate)
	http.HandleFunc("/projects/list", projects.ProjectList)
	http.HandleFunc("/projects/thankyou", projects.ProjectSaved)


	// Authentication
	// Github
	http.HandleFunc("/contributor/github/login/", authentication.GithubContributorLoginHandler)
	http.HandleFunc("/contributor/github/callback", authentication.GithubCallbackHandler)

	http.HandleFunc("/project/github/login/", authentication.GithubProjectLoginHandler)
	http.HandleFunc("/project/github/callback", authentication.GithubCallbackHandler)

	// Func to receive data after login
	http.HandleFunc("/github/loggedin", func(w http.ResponseWriter, r *http.Request) {
		authentication.GithubLoggedinHandler(w, r, "")
	})

    http.ListenAndServe(common.CONST_APP_PORT, nil)
}
