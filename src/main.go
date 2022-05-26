package main

import (
	"log"
	"net/http"
	"sources/app"
	"sources/authentication"

	"github.com/joho/godotenv"
)

type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

// init() executes before the main program
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
    fs := http.FileServer(http.Dir("../assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// Index route
	http.HandleFunc("/", app.IndexHandler)

    // Github Contributor Login route
	http.HandleFunc("/contributor/login/github/", authentication.GithubContributorLoginHandler)

	// Github Project Login route
	http.HandleFunc("/project/login/github/", authentication.GithubProjectLoginHandler)

	// Github Contributor callback
	http.HandleFunc("/contributor/github/callback", authentication.GithubCallbackHandler)

	// Github Project callback
	http.HandleFunc("/project/github/callback", authentication.GithubCallbackHandler)

	// Route where the authenticated user is redirected to
	http.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		authentication.GithubLoggedinHandler(w, r, "")
	})

    http.ListenAndServe(":8080", nil)
}

