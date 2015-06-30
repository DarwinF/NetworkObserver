package main

import (
	"html/template"
	"net/http"
)

func servePage(w http.ResponseWriter, r *http.Request, pageName string) {
	t, _ := template.ParseFiles(pageName)
	t.Execute(w, nil)
}

// This is the root webpaged served so you have to login everytime
func login(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "html/login.html")
}

// Verify if the user has loggid in properly
// TODO: replace 'success.html' with the dashboard when
// it has been created
func checkLogin(w http.ResponseWriter, r *http.Request) {

	//uname := r.FormValue("username")
	pword := r.FormValue("password")

	if pword != "password" {
		servePage(w, r, "html/error.html")
	} else {
		servePage(w, r, "html/dashboard.html")
	}
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/checkLogin", checkLogin)

	http.ListenAndServe(":8951", nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(":8951", "cert.pem", "key.pem", nil)
}
