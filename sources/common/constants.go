package common

const(
	CONST_APP_PORT = ":8080"

	CONST_GITHUB string = "github"

	// Session
	CONST_SESSION_NAME = "session"
	CONST_USER_NAME = "name"
	CONST_USER_IMAGE = "image"

	// User type
	CONST_USER_CONTRIBUTOR = "contributor"
	CONST_USER_PROJECT = "project"

	// DB Flags
	CONST_ACTIVE = 1
	CONST_INACTIVE = 0
	CONST_UNDER_MODERATION = 11

	// MongoDb Table names
	CONST_MO_USERS = "users"
	CONST_MO_USER_DETAILS = "user_details"
	CONST_MO_SOCIALS = "socials"
	CONST_MO_USER_SESSIONS = "user_sessions"
	CONST_MO_CONTRIBUTOR_PREFERENCES = "contributor_preferences"
	CONST_MO_PROJECTS = "projects"
	CONST_MO_CONTRIBUTOR_NOTIFICATIONS = "contributor_notifications"
	CONST_MO_PROJECT_NOTIFICATIONS = "project_notifications"
	CONST_MO_BOOKMARKS = "bookmarks"
	CONST_MO_REACTIONS = "reactions"
)
