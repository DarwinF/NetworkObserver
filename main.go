package main

import (
	"NetworkObserver/web"
	"net/http"
)

func main() {
	http.HandleFunc("/", webserv.Root)
	http.HandleFunc("/checkLogin", webserv.CheckLogin)
	http.HandleFunc("/dashboard", webserv.Dashboard)

	http.HandleFunc("/dashboard/settings", webserv.Settings)

	/*
		TODO: Create handle functions for these
		"/dashboard/configuration"
		"/dashboard/current_test"
		"/dashboard/results"
	*/

	http.ListenAndServe(":8951", nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(":8951", "cert.pem", "key.pem", nil)
}
