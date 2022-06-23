package contributors

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"techpro.club/sources/common"
	"techpro.club/sources/templates"
	"techpro.club/sources/templates/projects"
	"techpro.club/sources/users"
)
type FinalPreferencesOutStruct struct{
	ProgrammingLanguages map[string]string `json:"programmingLanguages"`
	AlliedServices map[string]string `json:"alliedServices"`
	ProjectType map[string]string `json:"projectType"`
	Contributors map[string]string `json:"contributors"`
	ContributorPreferences common.SaveContributorPreferencesStruct `json:"contributorPreferences"`
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
		"contains" : projects.Contains,
		"sliceToCsv" : projects.SliceToCsv,
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

	if r.Method == "GET"{
		preferences := fetchPreferences(userID)

		finalPreferencesOutStruct := FinalPreferencesOutStruct{
			ProgrammingLanguages,
			AlliedServices,
			ProjectType,
			Contributors,
			preferences,
		}

		tmpl, err := template.New("").Funcs(functions).ParseFiles("templates/app/contributors/preferences.gohtml", "templates/app/contributors/common/base.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "contributorbase", finalPreferencesOutStruct) 
		}
	} else {
	
		errParse := r.ParseForm()
		if errParse != nil {
			log.Println(errParse.Error())
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

			client, _ := common.Mongoconnect()
			defer client.Disconnect(context.TODO())
	
			dbName := common.GetMoDb()
			saveContributorPreference := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_PREFERENCES)
	
			_, err := saveContributorPreference.InsertOne(context.TODO(), result)
	
			if err != nil {
				fmt.Println(err)
			}
			
			

			http.Redirect(w, r, "/contributors/thankyou", http.StatusSeeOther)
		}
	}
}

// Return contributor preferences, if already saved
func fetchPreferences(userID string) (preferences common.SaveContributorPreferencesStruct){
	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchPreferences := client.Database(dbName).Collection(common.CONST_MO_CONTRIBUTOR_PREFERENCES)
	err := fetchPreferences.FindOne(context.TODO(),  bson.M{"userid": userID}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&preferences)

	if err != nil {
		fmt.Println(err, userID)
	} 

	return preferences
}