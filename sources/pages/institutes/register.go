package institutes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/users"
)

type InstitutePageStruct struct {
	InstituteData      common.FetchInstitutetruct      `json:"instituteData"`
	UserNameImage      common.UsernameImageStruct      `json:"userNameImage"`
	NotificaitonsCount int64                           `json:"notificationsCount"`
	NotificationsList  []common.MainNotificationStruct `json:"nofiticationsList"`
	PageDetails        common.PageDetails              `json:"pageDetails"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/institute/register" {
		pages.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if !sessionOk {

		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	// Fetch notificaitons
	_, _, notificationsCount, notificationsList := pages.NotificationsCountAndTopFive(userID)

	if !status {
		log.Println(msg)
	} else {
		userNameImage = common.UsernameImageStruct{Username: userName, Image: image}
	}

	_, _, instituteData := GetUnregisteredInstitute(userID)

	fmt.Println(instituteData)

	if r.Method == "POST" {
		// Update user profile

		errParse := r.ParseForm()

		if errParse != nil {
			log.Println(errParse.Error())
		} else {

			institute := common.SaveInstitutetruct{}

			institute.ImageLink = r.Form.Get("imageLink")
			institute.Name = r.Form.Get("name")
			institute.Vision = r.Form.Get("vision")
			institute.Mission = r.Form.Get("mission")
			institute.Founded = r.Form.Get("founded")
			institute.About = r.Form.Get("about")
			institute.Email = r.Form.Get("email")
			institute.Website = r.Form.Get("website")
			institute.Landline = r.Form.Get("landline")
			institute.Mobile = r.Form.Get("mobile")
			institute.Twitter = r.Form.Get("twitter")
			institute.Facebook = r.Form.Get("facebook")
			institute.LinkedIn = r.Form.Get("linkedin")
			institute.UserId = userID
			institute.VerifiedBy = instituteData.VerifiedBy
			institute.IsVerified = instituteData.IsVerified

			instituteStatus, instituteMsg := false, ""
			if instituteData.Name != "" {
				instituteStatus, instituteMsg = UpdateInstitute(institute, instituteData.Id)
			} else {
				instituteStatus, instituteMsg = SaveInstitute(institute)
			}

			if instituteStatus {
				fmt.Println("ok", instituteMsg)
			} else {
				fmt.Println("Wrong", instituteMsg)
			}
		}
	}

	_, _, instituteData = GetUnregisteredInstitute(userID)

	baseUrl := common.GetBaseurl() + common.CONST_APP_PORT
	pageDetails := common.PageDetails{BaseUrl: baseUrl, Title: "Register your institute"}

	output := InstitutePageStruct{
		UserNameImage:      userNameImage,
		NotificaitonsCount: notificationsCount,
		NotificationsList:  notificationsList,
		PageDetails:        pageDetails,
		InstituteData:      instituteData,
	}

	tmpl, err := template.New("").ParseFiles("templates/app/common/base.gohtml", "templates/app/common/contributormenu.gohtml", "templates/app/institutes/register.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		tmpl.ExecuteTemplate(w, "base", output)
	}

}
