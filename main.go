//--------------------------------------------
// main.go
//
// Initializes and sets up Network Observer
// for use.
//--------------------------------------------

package main

import (
	"NetworkObserver/configuration"
	"NetworkObserver/logger"
	"NetworkObserver/settings"
	"NetworkObserver/web"
	"net/http"
	"os"
)

// Random default value
var portNumber string = "5000"
var loc string = settings.AppLocation

func init() {
	// Check if the application directory exists, create it if it doesn't
	if _, err := os.Stat(loc); os.IsNotExist(err) {
		os.Mkdir(loc, 0600)
	}

	cl := loc + "/.cookies"
	logger.WriteString("Removing the existing .cookies file")
	os.Remove(cl)

	// this is here because apparently auth's init() is called before this one is
	if _, err := os.Stat(cl); os.IsNotExist(err) {
		logger.WriteString("Creating a new .cookies file.")
		file, _ := os.Create(cl)
		file.Close()
	}

	if _, err := os.Stat(loc + "/.password"); os.IsNotExist(err) {
		logger.WriteString("The .password file does not exist. Creating a new .password file.")
		file, _ := os.Create(loc + ".password")
		file.Close()
	}

	logger.WriteString("Setting the port number to " + configuration.GetPortNumber())
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
	http.HandleFunc("/logout", web.Logout)

	// Dashboard Pages
	http.HandleFunc("/dashboard/configure", web.Configure)
	http.HandleFunc("/dashboard/start_test", web.StartTest)
	http.HandleFunc("/dashboard/reports", web.Reports)

	// Start the Webserver
	http.ListenAndServe(portNumber, nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(portNumber, "cert.pem", "key.pem", nil)
}
