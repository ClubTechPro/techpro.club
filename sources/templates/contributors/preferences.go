package contributors

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"techpro.club/sources/common"
	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)
type FinalPreferencesOutStruct struct{
	ProgrammingLanguages map[string]string `json:"programmingLanguages"`
	AlliedServices map[string]string `json:"alliedServices"`
	ProjectType map[string]string `json:"projectType"`
	Contributors map[string]string `json:"contributors"`
	ContributorPreferences common.SaveContributorPreferencesStruct `json:"contributorPreferences"`
	UserNameImage common.UsernameImageStruct `json:"userNameImage"`
}

func Preferences(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/contributors/preferences" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
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

	var functions = template.FuncMap{
		"contains" : templates.Contains,
		"sliceToCsv" : templates.SliceToCsv,
	}

	ProgrammingLanguages := map[string]string{
		"c" : "C",
		"cpp" : "C++",
		"csharp" : "C#",
		"clojure" : "Clojure",
		"codeql" : "CodeQL",
		"coffeescript" : "CoffeeScript",
		"dm" : "DM",
		"dart" : "Dart",
		"elixir" : "Elixir",
		"elm" : "Elm",
		"emacslisp" : "Emacs Lisp",
		"erlang" : "Erlang",
		"fsharp" : "F#",
		"fortran" : "Fortran",
		"go" : "Golang",
		"groovy" : "Groovy",
		"java" : "Java",
		"js" : "Javascript",
		"jinja" : "Jinja",
		"jsonnet" : "Jsonnet",
		"julia" : "Julia",
		"kotlin" : "Kotlin",
		"lean" : "Lean",
		"lua" : "Lua",
		"matlab" : "Matlab",
		"nix" : "Nix",
		"objectivec" : "Objective-C",
		"ocaml" : "OCaml",
		"perl" : "Perl",
		"php" : "PHP",
		"powershell" : "Powershell",
		"puppet" : "Puppet",
		"python" : "Python",
		"r" : "R",
		"roff" : "Roff",
		"ruby" : "Ruby",
		"rust" : "Rust",
		"scala" : "Scala",
		"scss" : "SCSS",
		"shell" : "Shell",
		"swift" : "Swift",
		"systemverilog" : "System Verilog",
		"typescript" : "Typescript",
		"vala" : "Vala",
		"verilog" : "Verilog",
		"vimscript" : "Vim script",
		"vbnet" : "Visual Basic .NET",
		"wasm" : "WebAssembly",
		"yaml" : "YAML",
	}
	
	AlliedServices := map[string] string {
		"devops" : "DevOps",
		"documentation" : "Documentation",
		"sanitization" : "Code Sanitization",
		"test" : "Test Cases",
	}
	
	ProjectType := map[string]string{
		"library_plugin" : "Library/Plugin",
		"database" : "Database",
		"webapp" : "Web Application",
		"mobileapp" : "Mobile Application", 
		"desktopapp" : "Desktop Application",
		"others" : "Others",
	}

	Contributors := map[string]string{
		"1" : "Project founder only",
		"less_than_10" : "Less than 10",
		"more_than_10" : "More than 10",
	}

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := templates.FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}


	if r.Method == "GET"{
		_, _, preferences := fetchPreferences(userID)

		finalPreferencesOutStruct := FinalPreferencesOutStruct{
			ProgrammingLanguages,
			AlliedServices,
			ProjectType,
			Contributors,
			preferences,
			userNameImage,
		}

		tmpl, err := template.New("").Funcs(functions).ParseFiles("templates/app/contributors/preferences.gohtml", "templates/app/contributors/common/base_new.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "contributorbase", finalPreferencesOutStruct) 
		}
	} else {
	
		status, msg := savePreferences(w, r, userID)
		if !status{
			fmt.Println(msg)
		} else {
			http.Redirect(w, r, "/contributors/thankyou", http.StatusOK)
		}
	}
}

// Return contributor preferences, if already saved
func fetchPreferences(userID primitive.ObjectID) (status bool, msg string, preferences common.SaveContributorPreferencesStruct){
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchPreferences := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_PREFERENCES)
	results, err := fetchPreferences.Find(context.TODO(),  bson.M{"userid": userID}, options.Find().SetProjection(bson.M{"_id": 0}))

	if err != nil {
		msg = err.Error()
		status = false
	} else {
	
		for results.Next(context.TODO()) {

			errDecode := results.Decode(&preferences)

			if errDecode != nil {
				msg = errDecode.Error()
				status = false
			} else {
				msg= "Success"
				status = true
			}
		}
	}

	return status, msg, preferences
}

// Save preferences for contributor
func savePreferences(w http.ResponseWriter, r *http.Request, userID primitive.ObjectID) (status bool, msg string){
	status = false
	msg = ""


	errParse := r.ParseForm()
	if errParse != nil {
		msg = errParse.Error()
	} else {
		languages := r.Form["language"]
		otherLanguages := r.Form.Get("otherLanguages")
		allied := r.Form["allied"]
		notificationFrequency := r.Form.Get("emailFrequency")
		projectType := r.Form["pType"]
		contributorCount := r.Form.Get("contributorCount")
		paidJob :=  r.Form.Get("paidJob")
		relocation := r.Form.Get("relocation")
		qualification := r.Form.Get("qualification")


		otherLanguagesSplit := strings.Split(otherLanguages, ",")

		result := common.SaveContributorPreferencesStruct{userID, languages, otherLanguagesSplit, allied, projectType, notificationFrequency, contributorCount, paidJob, relocation, qualification}

		_, _, client := common.Mongoconnect()
		defer client.Disconnect(context.TODO())

		dbName := common.GetMoDb()
		saveContributorPreference := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_PREFERENCES)

		_, err := saveContributorPreference.InsertOne(context.TODO(), result)

		if err != nil {
			msg = err.Error()
		} else {
			status = true
			msg = "Success"
		}
	}

	return status, msg
}