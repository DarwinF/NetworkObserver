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
	http.HandleFunc("/", web.Root)
	http.HandleFunc("/checkLogin", web.CheckLogin)
	http.HandleFunc("/dashboard", web.Dashboard)
	http.HandleFunc("/createaccount", web.CreateAccount)
	http.HandleFunc("/account", web.HandleAccount)

	// Handlers
	http.HandleFunc("/saveConfig", web.SaveConfig)
	http.HandleFunc("/savetest", web.SaveTest)
	http.HandleFunc("/teststarted", web.TestStarted)

	// Dashboard Pages
	http.HandleFunc("/dashboard/configure", web.Configure)
	http.HandleFunc("/dashboard/start_test", web.StartTest)
	http.HandleFunc("/dashboard/reports", web.Reports)

	// Start the Webserver
	http.ListenAndServe(portNumber, nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(portNumber, "cert.pem", "key.pem", nil)
}
