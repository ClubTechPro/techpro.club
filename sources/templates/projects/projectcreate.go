package projects

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"techpro.club/sources/common"
	"techpro.club/sources/templates"
	"techpro.club/sources/users"
)

type FinalOutStruct struct{
	ProgrammingLanguages map[string]string `json:"programmingLanguages"`
	AlliedServices map[string]string `json:"alliedServices"`
	ProjectType map[string]string `json:"projectType"`
	Contributors map[string]string `json:"contributors"`
}

func ProjectCreate(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/projects/create" {
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

		constantLists := FinalOutStruct{
			ProgrammingLanguages,
			AlliedServices,
			ProjectType,
			Contributors,
		}

		tmpl, err := template.New("").ParseFiles("templates/app/projects/projectcreate.gohtml", "templates/app/projects/common/base.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}else {
			tmpl.ExecuteTemplate(w, "projectbase", constantLists) 
		}

	} else {
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
			var result common.SaveProjectStruct

			if submit == "Save as draft" {
				result = common.SaveProjectStruct{userID, projectName, projectDescription, repoLink, language, otherLanguagesSplit, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, dt, "", "", common.CONST_INACTIVE}
				saveProject(w, r, result)
			} else {
				result = common.SaveProjectStruct{userID, projectName, projectDescription, repoLink, language, otherLanguagesSplit, allied, projectType, contributorCount, documentation, public, company, companyName ,funded, dt, dt, "", common.CONST_UNDER_MODERATION}
				saveProject(w, r, result)
			}	
		}
	}
}


func saveProject(w http.ResponseWriter, r *http.Request, newProjectStruct common.SaveProjectStruct){
	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveProject := client.Database(dbName).Collection(common.CONST_PR_PROJECTS)

	_, err := saveProject.InsertOne(context.TODO(), newProjectStruct)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/projects/thankyou", http.StatusSeeOther)
}
