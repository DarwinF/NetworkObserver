package main

import (
	"NetworkObserver/web"
	"net/http"
)

func main() {
	// Base Pages
	http.HandleFunc("/", webserv.Root)
	http.HandleFunc("/checkLogin", webserv.CheckLogin)
	http.HandleFunc("/dashboard", webserv.Dashboard)

	// Dashboard Pages
	http.HandleFunc("/dashboard/settings", webserv.Settings)
	http.HandleFunc("/dashboard/configuration", webserv.Configure)
	http.HandleFunc("/dashboard/current_test", webserv.CurrentTest)
	http.HandleFunc("/dashboard/results", webserv.Results)

	http.ListenAndServe(":8951", nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(":8951", "cert.pem", "key.pem", nil)
}
