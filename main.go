//--------------------------------------------
// main.go
//
// Initializes and sets up Network Observer
// for use.
//--------------------------------------------

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
	http.HandleFunc("/dashboard/configure", webserv.Configure)
	http.HandleFunc("/dashboard/start_test", webserv.StartTest)
	http.HandleFunc("/dashboard/reports", webserv.Reports)

	// Start the Webserver
	http.ListenAndServe(":8951", nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(":8951", "cert.pem", "key.pem", nil)
}
