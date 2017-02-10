package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ScriptIndex",
		"GET",
		"/scripts",
		ScriptIndex,
	},
	Route{
		"ScriptShow",
		"GET",
		"/script/{scriptId}",
		ScriptShow,
	},
	Route{
		"ScriptDelete",
		"DELETE",
		"/script/{scriptId}",
		ScriptDelete,
	},
	Route{
		"ScriptCheckin",
		"POST",
		"/script/checkin/{scriptName}",
		ScriptCheckin,
	},
	Route{
		"ScriptSchedule",
		"POST",
		"/script/{scriptId}/schedule",
		ScriptSchedule,
	},
	Route{
		"MailIndex",
		"GET",
		"/mail/",
		MailIndex,
	},
	Route{
		"MailCreate",
		"POST",
		"/mail/",
		MailCreate,
	},
	Route{
		"MailDelete",
		"DELETE",
		"/mail/{mailId}/delete",
		MailDelete,
	},
}
