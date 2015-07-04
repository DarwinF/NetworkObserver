package main

import (
	"html/template"
	"net/http"
	"NetworkObserver/auth"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/checkLogin", checkLogin)
	http.HandleFunc("/dashboard", dashboard)

	http.ListenAndServe(":8951", nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(":8951", "cert.pem", "key.pem", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "html/login.html")
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePage(w, r, "html/dashboard.html")
	} else {
		servePage(w, r, "html/login.html")
	}
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
		servePage(w, r, "html/dashboard.html")
	} else {
		servePage(w, r, "html/error.html")
	}
}