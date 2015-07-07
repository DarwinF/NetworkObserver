package webserv

import (
	"NetworkObserver/auth"
	"html/template"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		servePage(w, r, "html/login.html")
	}
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, "html/dashboard.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	uname := r.FormValue("username")
	pword := r.FormValue("password")

	// Authenticate user credentials
	authenticated := auth.CheckCredentials(uname, pword)

	if authenticated == true {
		auth.SetSessionID(w)
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		servePage(w, r, "html/error.html")
	}
}

func Settings(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, "htm/dashboard/settings.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Configure(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, "htm/dashboard/config.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func CurrentTest(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, "htm/dashboard/test.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Results(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, "htm/dashboard/results.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func servePage(w http.ResponseWriter, r *http.Request, pageName string) {
	t, _ := template.ParseFiles(pageName)
	t.Execute(w, nil)
}
