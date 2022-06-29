package templates

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"techpro.club/sources/common"
	"techpro.club/sources/users"
)

type SettingsStruct struct{
	UserProfile common.FetchUserStruct `json:"userprofile"`
	UserSocials common.FetchSocialStruct `json:"socials"`
}

// Display and edit user profile
func UserSettings(w http.ResponseWriter, r *http.Request) {

	// Check if the path is ok
	// Check if the session is ok

	// Fetch name and image from cookies
	// Fetch email, _id, location,repourl from database
	// Display the user's profile
	// Display the user's settings

	
	
	if r.URL.Path != "/users/settings" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} 


	if r.Method == "GET"{
		userprofile := fetchUserProfile(userID)
		socials := fetchSocials(userID)

		userSettingsStruct := SettingsStruct{userprofile, socials}
		fmt.Println(userSettingsStruct)

		tmpl, err := template.New("").ParseFiles("templates/app/settings.gohtml", "templates/app/contributors/common/base.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "contributorbase", userSettingsStruct) 
		}

	} else {
		errParse := r.ParseForm()
		if errParse != nil {
			log.Println(errParse.Error())
		} else {
			name := r.Form.Get("name")
			repoUrl := r.Form.Get("otherLanguages")
			facebook := r.Form.Get("facebook")
			linkedin := r.Form.Get("linkedin")
			twitter := r.Form.Get("twitter")
			stackoverflow := r.Form.Get("stackoverflow")
			
			fmt.Println(name, repoUrl, facebook, linkedin, twitter, stackoverflow)
		}
	}
}

// Fetch user profile
func fetchUserProfile(userID string)(userProfile common.FetchUserStruct){
	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	userIDHex, _ := primitive.ObjectIDFromHex(userID)

	dbName := common.GetMoDb()
	fetchPreferences := client.Database(dbName).Collection(common.CONST_MO_USERS)
	err := fetchPreferences.FindOne(context.TODO(),  bson.M{"_id": userIDHex}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&userProfile)

	if err != nil {
		fmt.Println(err, userID)
	} 

	return userProfile
}

// Update user profile
func UpdateUserProfile(){

}

// Fetch socials
func fetchSocials(userID string)(socials common.FetchSocialStruct){
	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchPreferences := client.Database(dbName).Collection(common.CONST_MO_SOCIALS)
	err := fetchPreferences.FindOne(context.TODO(),  bson.M{"userid": userID}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&socials)

	if err != nil {
		fmt.Println(err, userID)
	} 

	fmt.Println(socials)
	return socials
}

// Update socials
func updateSocials(){

}