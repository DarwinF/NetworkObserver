package main

import (
	"NetworkObserver/auth"
	"html/template"
	"net/http"
)

//----------------------------------------------
// Variables
//----------------------------------------------
var previousPage string = ""

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/checkLogin", checkLogin)

	http.HandleFunc("/settings", softwareSettings)
	http.HandleFunc("/configure", testConfig)
	http.HandleFunc("/test", currTest)
	http.HandleFunc("/results", testResults)

	http.ListenAndServe(":8951", nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(":8951", "cert.pem", "key.pem", nil)
}

func servePage(w http.ResponseWriter, r *http.Request, pageName string) {
	t, _ := template.ParseFiles(pageName)
	t.Execute(w, nil)
}

func checkLogin(w http.ResponseWriter, r *http.Request) {
	uname := r.FormValue("username")
	pword := r.FormValue("password")

	// Authenticate user credentials
	authenticated := auth.CheckCredentials(uname, pword)

	if authenticated == true {
		auth.SetSessionID(w)
		redirect(w, r)
	} else {
		servePage(w, r, "html/error.html")
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	if previousPage == "" {
		servePage(w, r, "html/dash-redirect.html")
	} else {
		servePage(w, r, previousPage)
	}
}

func invalidSessionID(callingPage string, w http.ResponseWriter, r *http.Request) {
	previousPage = callingPage
	servePage(w, r, "html/login.html")
}

func root(w http.ResponseWriter, r *http.Request) {
	page := "html/dashboard.html"
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, page)
	} else {
		invalidSessionID(page, w, r)
	}
}

func softwareSettings(w http.ResponseWriter, r *http.Request) {
	page := "html/dashboard/settings.html"
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, page)
	} else {
		invalidSessionID(page, w, r)
	}
}

func testConfig(w http.ResponseWriter, r *http.Request) {
	page := "html/dashboard/config.html"
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, page)
	} else {
		invalidSessionID(page, w, r)
	}
}

func currTest(w http.ResponseWriter, r *http.Request) {
	page := "html/dashboard/test.html"
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, page)
	} else {
		invalidSessionID(page, w, r)
	}
}

func testResults(w http.ResponseWriter, r *http.Request) {
	page := "html/dashboard/results.html"
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, page)
	} else {
		invalidSessionID(page, w, r)
	}
}
