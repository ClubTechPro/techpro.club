package projects

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)

type FinalProjectOutStruct struct{
	ProgrammingLanguages map[string]string `json:"programmingLanguages"`
	AlliedServices map[string]string `json:"alliedServices"`
	ProjectType map[string]string `json:"projectType"`
	Contributors map[string]string `json:"contributors"`
	ProjectStruct ProjectStruct `json:"projectStruct"`
}

func ProjectEdit(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/projects/edit" {
        templates.ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/projects", http.StatusSeeOther)
	}

	var functions = template.FuncMap{
		"contains" : contains,
		"sliceToCsv" : sliceToCsv,
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

		projectID := r.URL.Query().Get("projectid")

		result := FetchProjectDetails(projectID, userID)

		constantLists := FinalProjectOutStruct{
			ProgrammingLanguages,
			AlliedServices,
			ProjectType,
			Contributors,
			result,
		}
		

		tmpl, err := template.New("").Funcs(functions).ParseFiles("templates/app/projects/projectedit.gohtml", "templates/app/projects/common/base.html")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "projectbase", constantLists) 
		}
	} else {
		projectID := r.URL.Query().Get("projectid")

		errParse := r.ParseForm()
		if errParse != nil {
			fmt.Println(errParse.Error())
		} else {
			projectName := r.Form.Get("projectName")
			repoLink := r.Form.Get("repoLink")
			projectDescription := r.Form.Get("projectDescription")
			language := r.Form["language"]
			otherLanguages := r.Form.Get("otherLanguages")
			allied := r.Form["allied"]
			projectType :=  r.Form["pType"]
			contributorCount := r.Form.Get("contributorCount")
			documentation := r.Form.Get("documentation")
			public := r.Form.Get("public")
			company := r.Form.Get("company")
			companyName := r.Form.Get("companyName")
			funded := r.Form.Get("funded")
			submit := r.Form.Get("submit")

			otherLanguagesSplit := strings.Split(otherLanguages, ",")

			timeNow := time.Now()
			dt := timeNow.Format(time.UnixDate)
			var result NewProjectStruct

			if submit == "Save as draft" {
				result = NewProjectStruct{userID, projectName, projectDescription, repoLink, language, otherLanguagesSplit, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, dt, "", "", common.CONST_INACTIVE}
				updateProject(w, r, projectID, result)
			} else {
				result = NewProjectStruct{userID, projectName, projectDescription, repoLink, language, otherLanguagesSplit, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, dt, dt, "", common.CONST_UNDER_MODERATION}
				updateProject(w, r, projectID, result)
			}	
		}
	}
}

func FetchProjectDetails(projectID, userID string) (projectDetails ProjectStruct){
	if(projectID != ""){

		projectIdHex, err := primitive.ObjectIDFromHex(projectID)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			client, _ := common.Mongoconnect()
			defer client.Disconnect(context.TODO())

			dbName := common.GetMoDb()
			fetchProject := client.Database(dbName).Collection(common.CONST_PR_PROJECTS)
			err := fetchProject.FindOne(context.TODO(),  bson.M{"userid": userID, "_id": projectIdHex}).Decode(&projectDetails)

			if err != nil {
				fmt.Println(err)
			} 
		}
	}

	return projectDetails
}

func updateProject(w http.ResponseWriter, r *http.Request, projectID string, newProjectStruct NewProjectStruct){
	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveProject := client.Database(dbName).Collection(common.CONST_PR_PROJECTS)

	projectIdHex, _ := primitive.ObjectIDFromHex(projectID)
	_, err := saveProject.UpdateOne(context.TODO(), bson.M{"_id": projectIdHex}, bson.M{"$set": newProjectStruct})

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/projects/thankyou", http.StatusSeeOther)
}

// Check if a string exists in a slice.
func contains(s []string, e string) (status bool) {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

// Convert slice of strings to csv string
func sliceToCsv(s []string) (csv string){
	csv = strings.Join(s, ",")
	return csv
}