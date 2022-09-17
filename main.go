package main

import (
	"log"
	"net/http"

	"techpro.club/sources/authentication"
	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/pages/contributors"
	"techpro.club/sources/pages/projects"
	"techpro.club/sources/pages/videos"

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

	// APIs
	http.HandleFunc("/api/managereaction", pages.ManageReactions)
	http.HandleFunc("/api/managebookmark", pages.ManageBookmarks)
	http.HandleFunc("/api/marknotificationsread", pages.MarkNotificationRead)
	http.HandleFunc("/api/deleteuser", pages.DeleteUser)
	http.HandleFunc("/api/deleteproject", projects.DeleteProject)

	// Templates
	http.HandleFunc("/", pages.IndexHandler)
	http.HandleFunc("/contactus", pages.ContactUs)
	http.HandleFunc("/careers", pages.Careers)
	http.HandleFunc("/company", pages.Company)
	http.HandleFunc("/brand", pages.Brand)
	http.HandleFunc("/campus", pages.Campus)
	http.HandleFunc("/campusonboard", pages.CampusOnboard)
	http.HandleFunc("/videos", pages.Videos)
	http.HandleFunc("/privacy-policy", pages.PrivacyPolicy)
	http.HandleFunc("/cookie-policy", pages.CookiePolicy)
	http.HandleFunc("/terms-and-conditions", pages.TermsOfService)

	// Users
	http.HandleFunc("/users/editprofile", pages.UserEdit)
	http.HandleFunc("/users/profiles", pages.PublicProfile)
	http.HandleFunc("/users/profile", pages.Profile)
	http.HandleFunc("/users/notifications", pages.Notifications)
	http.HandleFunc("/users/settings", pages.UserSettings)

	// Templates/Contributors
	http.HandleFunc("/contributors/feeds", contributors.Feeds)
	http.HandleFunc("/contributors/videofeeds", contributors.VideoFeeds)
	http.HandleFunc("/contributors/preferences", contributors.Preferences)
	http.HandleFunc("/contributors/thankyou", contributors.PreferencesSaved)
	http.HandleFunc("/contributors/reactions", contributors.FetchReactions)
	http.HandleFunc("/contributors/bookmarks", contributors.FetchBookmarks)
	

	// Templates/Projects
	http.HandleFunc("/projects/create", projects.ProjectCreate)
	http.HandleFunc("/projects/list", projects.ProjectList)
	http.HandleFunc("/projects/view", projects.ProjectPreview)
	http.HandleFunc("/projects/edit", projects.ProjectEdit)
	http.HandleFunc("/projects/thankyou", projects.ProjectSaved)

	// Templates/Videos
	http.HandleFunc("/videos/list", videos.VideosList)
	http.HandleFunc("/videos/newvideo", videos.NewVideo)
	http.HandleFunc("/videos/editvideo", videos.EditVideo)



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

	http.HandleFunc("/logout", pages.Logout)

	// Start the web server
    http.ListenAndServe(common.CONST_APP_PORT, nil)
}
