package projects

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/users"
)

func DeleteProject(w http.ResponseWriter, r *http.Request){

	status := false
	msg := ""

	if r.URL.Path != "/api/deleteproject" {
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

	projectID := r.URL.Query().Get("projectid")

	_, _,  client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchProjects := client.Database(dbName).Collection(common.CONST_MO_PROJECTS)

	_, err:= fetchProjects.DeleteOne(context.TODO(), bson.M{"_id": pages.StringToObjectId(projectID), "userid": userID})

	if err != nil{
		msg = err.Error()
	} else {
		msg = "Success"
		status = true
	}

	output := common.JsonOutput{
		Status: status,
		Msg: msg,
	}

	out, _ := json.Marshal(output)
	w.Write(out)
}