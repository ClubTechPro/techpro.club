package contributors

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/users"
)

type FinalProjectBookmarkListOutStruct struct{
	ProjectsList []common.FetchProjectStruct `json:"projectsList"`
	UserNameImage common.UsernameImageStruct `json:"userNameImage"`
}

func Bookmarks(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/contributors/bookmarks" {
        pages.ErrorHandler(w, r, http.StatusNotFound)
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

	var finalOutStruct FinalProjectBookmarkListOutStruct
	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}
	
	_, _, results := listBookmarks(w, r, userID)
	finalOutStruct = FinalProjectBookmarkListOutStruct{results, userNameImage}
	

	tmpl, err := template.New("").ParseFiles("templates/app/projects/projectlist.gohtml", "templates/app/projects/common/base_new.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "projectbase", finalOutStruct) 
	}
}

// List bookmarks for a user
func listBookmarks(w http.ResponseWriter, r *http.Request, userID primitive.ObjectID)(status bool, msg string, results []common.FetchProjectStruct){
	
	status = false
	msg = ""
	var bookmarkResult common.BookmarkStruct

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchBookmarks := client.Database(dbName).Collection(common.CONST_MO_BOOKMARKS)
	errBookmarks := fetchBookmarks.FindOne(context.TODO(),  bson.M{"userid": userID}).Decode(&bookmarkResult)


	if errBookmarks != nil {
		msg = errBookmarks.Error()
	} else {
		fetchProjects := client.Database(dbName).Collection(common.CONST_MO_BOOKMARKS)
		projectsList, errProjectsList := fetchProjects.Find(context.TODO(),  bson.M{"_id": bson.M{"$in": bookmarkResult} })

		if errProjectsList != nil {
			msg = errProjectsList.Error()
		} else {
			for projectsList.Next(context.TODO()) {
				var result common.FetchProjectStruct
				errProjectsList := projectsList.Decode(&result)
				if errProjectsList != nil {
					msg = errProjectsList.Error()
				} else {
					results = append(results, result)
					status = true
					msg = "Success"
					fmt.Println(result)
				}
			}
		}
	}

	return status, msg, results
}

// Add a bookmark for a user
func AddBookmark(userID primitive.ObjectID, projectID primitive.ObjectID)(status bool, msg string){
	
	status = false
	msg = ""
	var bookmarkResult common.BookmarkStruct

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchBookmarks := client.Database(dbName).Collection(common.CONST_MO_BOOKMARKS)
	errBookmarks := fetchBookmarks.FindOne(context.TODO(),  bson.M{"userid": userID}).Decode(&bookmarkResult)

	if errBookmarks != nil {
		fetchBookmarks.InsertOne(context.TODO(), bson.M{"userid": userID, "projects": projectID})
		status = true
		msg = "Success"
	} else {
		fetchBookmarks.UpdateOne(context.TODO(), bson.M{"userid": userID}, bson.M{"$addToSet": bson.M{"projects": projectID}})
		status = true
		msg = "Success"
	}

	return status, msg
}

// Remove a bookmark for a user
func RemoveBookmark(userID primitive.ObjectID, projectID primitive.ObjectID)(status bool, msg string){
	
	status = false
	msg = ""
	var bookmarkResult common.BookmarkStruct

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchBookmarks := client.Database(dbName).Collection(common.CONST_MO_BOOKMARKS)
	errBookmarks := fetchBookmarks.FindOne(context.TODO(),  bson.M{"userid": userID}).Decode(&bookmarkResult)

	if errBookmarks != nil {
		msg = errBookmarks.Error()
	} else {
		fetchBookmarks.UpdateOne(context.TODO(), bson.M{"userid": userID}, bson.M{"$pull": bson.M{"projects": projectID}})
		status = true
		msg = "Success"
	}

	return status, msg
}