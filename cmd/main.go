//--------------------------------------------
// main.go
//
// Initializes and sets up Network Observer
// for use.
//--------------------------------------------

package main

import (
	"net/http"
	"os"

	logger "github.com/darwinfroese/networkobserver/pkg/logging"
	settings "github.com/darwinfroese/networkobserver/pkg/settings"
	websrv "github.com/darwinfroese/networkobserver/pkg/webserver"

	"github.com/darwinfroese/networkobserver/pkg/configuration"
)

// Random default value
var portNumber = "5000"
var loc = settings.AppLocation

func init() {
	cl := loc + "/.cookies"
	logger.WriteString("Removing the existing .cookies file")
	os.Remove(cl)

	// this is here because apparently auth's init() is called before this one is
	if _, err := os.Stat(cl); os.IsNotExist(err) {
		logger.WriteString("Creating a new .cookies file.")
		file, e := os.Create(cl)
		if e != nil {
			logger.WriteString("There was an error creating the .cookies file.")
		}
		file.Close()
	}

	if _, err := os.Stat(loc + "/.password"); os.IsNotExist(err) {
		logger.WriteString("The .password file does not exist. Creating a new .password file.")
		file, e := os.Create(loc + "/.password")
		if e != nil {
			logger.WriteString("There was an error creating the .password file.")
		}
		file.Close()
	}

	logger.WriteString("Setting the port number to " + configuration.GetPortNumber())
	portNumber = ":" + configuration.GetPortNumber()
}

func main() {
	// Base Pages
	http.HandleFunc("/", websrv.Root)
	http.HandleFunc("/checkLogin", websrv.CheckLogin)
	http.HandleFunc("/dashboard", websrv.Dashboard)
	http.HandleFunc("/createaccount", websrv.CreateAccount)
	http.HandleFunc("/account", websrv.HandleAccount)

	// Handlers
	http.HandleFunc("/saveConfig", websrv.SaveConfig)
	http.HandleFunc("/savetest", websrv.SaveTest)
	http.HandleFunc("/teststarted", websrv.TestStarted)
	http.HandleFunc("/logout", websrv.Logout)

	// Dashboard Pages
	http.HandleFunc("/dashboard/configure", websrv.Configure)
	http.HandleFunc("/dashboard/start_test", websrv.StartTest)
	http.HandleFunc("/dashboard/reports", websrv.Reports)

	// Start the Webserver
	http.ListenAndServe(portNumber, nil)
	// Enable SSL and HTTPS connections
	//http.ListenAndServeTLS(portNumber, "cert.pem", "key.pem", nil)
}
