package common

import "go.mongodb.org/mongo-driver/bson/primitive"

// Page title struct
type PageTitle struct {
	Title string `json:"title"`
}

// Feeds struct
type FeedStruct struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectName string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
	RepoLink string `json:"repoLink"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	Company string `json:"company"`
	CompanyName string `json:"companyName"`
	CreatedDate string `json:"createdDate"`
	Userdetails []FeedParentUserStruct `json:"userdetails"`
}


// Feed parent user struct
type FeedParentUserStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
	ImageLink string `json:"imageLink"`
}

// Fetch Notifications Struct
type FetchNotificationStruct struct {
	Id primitive.ObjectID `json:"id"`
	UserID primitive.ObjectID `json:"userID"`
	NotificationType string `json:"notificationType"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	Link string `json:"link"`
	CreatedDate string `json:"createdDate"`
	Read bool `json:"read"`
}

// Save Notifications Struct
type SaveNotificationStruct struct {
	UserID primitive.ObjectID `json:"userID"`
	NotificationType string `json:"notificationType"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	Link string `json:"link"`
	CreatedDate string `json:"createdDate"`
	Read bool `json:"read"`
}


// Fetch projects collection struct
type FetchProjectStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"userId"`
	ProjectName string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
	RepoLink string `json:"repoLink"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	ProjectType []string `json:"projectType"`
	ContributorCount string `json:"contributorCount"`
	Documentation string `json:"documentation"`
	Public string `json:"public"`
	Company string `json:"company"`
	CompanyName string `json:"companyName"`
	Funded string `json:"funded"`
	CreatedDate string `json:"createdDate"`
	PublishedDate string `json:"publishedDate"`
	ClosedDate string `json:"closedDate"`
	IsActive int `json:"isActive"`
}

// Save projects collection struct
type SaveProjectStruct struct{
	UserID primitive.ObjectID `json:"userId"`
	ProjectName string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
	RepoLink string `json:"repoLink"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	ProjectType []string `json:"projectType"`
	ContributorCount string `json:"contributorCount"`
	Documentation string `json:"documentation"`
	Public string `json:"public"`
	Company string `json:"company"`
	CompanyName string `json:"companyName"`
	Funded string `json:"funded"`
	CreatedDate string `json:"createdDate"`
	PublishedDate string `json:"publishedDate"`
	ClosedDate string `json:"closedDate"`
	IsActive int `json:"isActive"`
}

// Fetch users struct
type FetchUserStruct struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email 		string `json:"email"`
	Name 		string `json:"name"`
	Location 	string `json:"location"`
	ImageLink 	string `json:"imageLink"`
	RepoUrl 	string `json:"repoUrl"`
	Source 		string `json:"source"`
	CreatedDate string `json:"createdDate"`
}

// Save users struct
type SaveUserStruct struct {
	Email 		string `json:"email"`
	Name 		string `json:"name"`
	Location 	string `json:"location"`
	ImageLink 	string `json:"imageLink"`
	RepoUrl 	string `json:"repoUrl"`
	Source 		string `json:"source"`
	CreatedDate string `json:"createdDate"`
}

// Fetch user sessions Struct
type FetchUserSessionStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"userId"`
	SessionID string `json:"sessionId"`
}

// Save user sessions Struct
type SaveUserSessionStruct struct{
	UserID primitive.ObjectID `json:"userId"`
	SessionID string `json:"sessionId"`
}

// Fetch social Struct
type FetchSocialStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"userId"`
	Twitter string `json:"twitter"`
	Facebook string `json:"facebook"`
	LinkedIn string `json:"linkedin"`
	Stackoverflow string `json:"stackoverflow"`
}

// save social Struct
type SaveSocialStruct struct{
	UserID primitive.ObjectID `json:"userId"`
	Twitter string `json:"twitter"`
	Facebook string `json:"facebook"`
	LinkedIn string `json:"linkedin"`
	Stackoverflow string `json:"stackoverflow"`
}

// Fetch contributor preferences Struct
type FetchContributorPreferencesStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"userId"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	ProjectType []string `json:"projectType"`
	NotificationFrequency string `json:"notificationFrequency"`
	ContributorCount string `json:"contributorCount"`
	PaidJob string `json:"paidJob"`
	Relocation string `json:"relocation"`
	Qualification string `json:"qualification"`
}

// Save contributor preferences Struct
type SaveContributorPreferencesStruct struct{
	UserID primitive.ObjectID `json:"userId"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	ProjectType []string `json:"projectType"`
	NotificationFrequency string `json:"notificationFrequency"`
	ContributorCount string `json:"contributorCount"`
	PaidJob string `json:"paidJob"`
	Relocation string `json:"relocation"`
	Qualification string `json:"qualification"`
}

// Username and image struct
type UsernameImageStruct struct{
	Username string `json:"username"`
	Image string `json:"image"`
}

// Bookmark struct
type BookmarkStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectBookmarks []primitive.ObjectID `json:"projectBookmarks"`
}