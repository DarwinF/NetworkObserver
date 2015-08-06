//--------------------------------------------
// main.go
//
// Initializes and sets up Network Observer
// for use.
//--------------------------------------------

package main

import (
	"NetworkObserver/configuration"
	"NetworkObserver/web"
	"net/http"
)

// Random default value
var portNumber string = "5000"

func init() {
	// Check for configuration file
	portNumber = ":" + configuration.GetPortNumber()
}

func main() {
	// Base Pages
	http.HandleFunc("/", webserv.Root)
	http.HandleFunc("/checkLogin", webserv.CheckLogin)
	http.HandleFunc("/dashboard", webserv.Dashboard)
	http.HandleFunc("/createaccount", webserv.CreateAccount)
	http.HandleFunc("/account", webserv.HandleAccount)

	// Handlers
	http.HandleFunc("/saveConfig", webserv.SaveConfig)
	http.HandleFunc("/savetest", webserv.SaveTest)

	// Dashboard Pages
	http.HandleFunc("/dashboard/configure", webserv.Configure)
	http.HandleFunc("/dashboard/start_test", webserv.StartTest)
	http.HandleFunc("/dashboard/reports", webserv.Reports)

	// Start the Webserver
	http.ListenAndServe(portNumber, nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(portNumber, "cert.pem", "key.pem", nil)
}
