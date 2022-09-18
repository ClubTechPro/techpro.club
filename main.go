package main

import (
	"log"
	"net/http"
	"time"

	"techpro.club/sources/authentication"
	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/pages/contributors"
	"techpro.club/sources/pages/projects"
	"techpro.club/sources/pages/videos"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
    

	goMux := mux.NewRouter()

	// fs := http.FileServer(http.Dir("assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets", fs))
	goMux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	// APIs
	goMux.HandleFunc("/api/managereaction", pages.ManageReactions)
	goMux.HandleFunc("/api/managebookmark", pages.ManageBookmarks)
	goMux.HandleFunc("/api/marknotificationsread", pages.MarkNotificationRead)
	goMux.HandleFunc("/api/deleteuser", pages.DeleteUser)
	goMux.HandleFunc("/api/deleteproject", projects.DeleteProject)

	// Templates
	goMux.HandleFunc("/", pages.IndexHandler)
	goMux.HandleFunc("/contactus", pages.ContactUs)
	goMux.HandleFunc("/careers", pages.Careers)
	goMux.HandleFunc("/company", pages.Company)
	goMux.HandleFunc("/brand", pages.Brand)
	goMux.HandleFunc("/campus", pages.Campus)
	goMux.HandleFunc("/campusonboard", pages.CampusOnboard)
	goMux.HandleFunc("/videos", pages.Videos)
	goMux.HandleFunc("/privacy-policy", pages.PrivacyPolicy)
	goMux.HandleFunc("/cookie-policy", pages.CookiePolicy)
	goMux.HandleFunc("/terms-and-conditions", pages.TermsOfService)

	// Users
	goMux.HandleFunc("/users/editprofile", pages.UserEdit)
	goMux.HandleFunc("/users/profiles", pages.PublicProfile)
	goMux.HandleFunc("/users/profile", pages.Profile)
	goMux.HandleFunc("/users/notifications", pages.Notifications)
	goMux.HandleFunc("/users/settings", pages.UserSettings)
	goMux.HandleFunc("/users/profiletest/{username}", pages.ProfileTest)

	// Templates/Contributors
	goMux.HandleFunc("/contributors/feeds", contributors.Feeds)
	goMux.HandleFunc("/contributors/videofeeds", contributors.VideoFeeds)
	goMux.HandleFunc("/contributors/preferences", contributors.Preferences)
	goMux.HandleFunc("/contributors/thankyou", contributors.PreferencesSaved)
	goMux.HandleFunc("/contributors/reactions", contributors.FetchReactions)
	goMux.HandleFunc("/contributors/bookmarks", contributors.FetchBookmarks)
	

	// Templates/Projects
	goMux.HandleFunc("/projects/create", projects.ProjectCreate)
	goMux.HandleFunc("/projects/list", projects.ProjectList)
	goMux.HandleFunc("/projects/view", projects.ProjectPreview)
	goMux.HandleFunc("/projects/edit", projects.ProjectEdit)
	goMux.HandleFunc("/projects/thankyou", projects.ProjectSaved)

	// Templates/Videos
	goMux.HandleFunc("/videos/list", videos.VideosList)
	goMux.HandleFunc("/videos/newvideo", videos.NewVideo)
	goMux.HandleFunc("/videos/editvideo", videos.EditVideo)



	// Authentication
	// Github
	goMux.HandleFunc("/contributors/github/login/", authentication.GithubContributorLoginHandler)
	goMux.HandleFunc("/contributors/github/callback", authentication.GithubContributorCallbackHandler)

	goMux.HandleFunc("/projects/github/login/", authentication.GithubProjectLoginHandler)
	goMux.HandleFunc("/projects/github/callback", authentication.GithubProjectCallbackHandler)

	// Func to receive data after login
	goMux.HandleFunc("/github/loggedin", func(w http.ResponseWriter, r *http.Request) {
		authentication.GithubLoggedinHandler(w, r, "", "", "")
	})

	goMux.HandleFunc("/logout", pages.Logout)

	srv := &http.Server{
		Handler: goMux,
		Addr:    "127.0.0.1" + common.CONST_APP_PORT,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the web server
    srv.ListenAndServe()
}
