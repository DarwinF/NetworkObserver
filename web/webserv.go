//--------------------------------------------
// web/webserv.go
//
// Handles serving and authenticating for all
// the webpages.
//
// All the handler functions are declared in
// this file.
//--------------------------------------------

package webserv

import (
	"NetworkObserver/auth"
	"NetworkObserver/configuration"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

//--------------------------------------------
// Structs for Pages
//--------------------------------------------
type configPage struct {
	Port          string
	InternalAddr  string
	ExternalAddr  string
	ExternalURL   string
	SpeedFileLoc  string
	ReportFileLoc string
	PingDelay     string
	SpeedDelay    string
}

// All URLs default to this function
func Root(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		servePageStatic(w, r, "html/login.html")
	}
}

// Handles URLs referencing dashboard
func Dashboard(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePageStatic(w, r, "html/dashboard.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// Validates a login attempt or redirects the user to
// an error page which redirects them to "Root"
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	uname := r.FormValue("username")
	pword := r.FormValue("password")

	// Authenticate user credentials
	authenticated := auth.CheckCredentials(uname, pword)

	if authenticated == true {
		auth.SetSessionID(w)
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		servePageStatic(w, r, "html/error.html")
	}
}

// Create a new account by comparing both of the passwords
// entered and then hashing and storing the password
func CreateAccount(w http.ResponseWriter, r *http.Request) {

}

// Serve the webpage for creating an account
func Account(w http.ResponseWriter, r *http.Request) {

}

// Save the configuration settings and then reload the configuration
// page (which will automatically reload the new settings)
func SaveConfig(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		// Save Configuration Settings
		saveConfigToStruct(r)
		http.Redirect(w, r, "/dashboard/configure", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//--------------------------------------------
// Dashboard page handler functions
// The following four functions serve dynamic
// pages needed for the dashboard.
//
// Note: Currently static pages, will be dynamic
// later.
//--------------------------------------------
func Settings(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePageStatic(w, r, "html/dashboard/settings.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Configure(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	// Build struct to serve page with:
	configStruct := configPage{}
	buildConfigStruct(&configStruct)

	if valid == true {
		servePageDynamic(w, r, "html/dashboard/config.html", configStruct)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func StartTest(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	fmt.Println(r.FormValue("location"))

	if valid == true {
		servePageStatic(w, r, "html/dashboard/starttest.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Reports(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePageStatic(w, r, "html/dashboard/reports.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// Serves a static page
func servePageStatic(w http.ResponseWriter, r *http.Request, pageName string) {
	t, _ := template.ParseFiles(pageName)
	t.Execute(w, nil)
}

// Serves a page after gathering data needed
func servePageDynamic(w http.ResponseWriter, r *http.Request, pageName string, data interface{}) {
	t, _ := template.ParseFiles(pageName)
	t.Execute(w, data)
}

//----------------------------------------
// Configuration Handling Functions
//----------------------------------------

// Saves the setting sto the config file
func saveConfigToStruct(r *http.Request) {
	configuration.SetInternalIP(ipToMap(r.FormValue("internalip")))
	configuration.SetReportLocations(r.FormValue("reportfileloc"))
	configuration.SetPortNumber(r.FormValue("portnumber"))
	configuration.SetExternalIPs(lineToSlice(r.FormValue("externalip")))
	configuration.SetExternalURLs(lineToSlice(r.FormValue("externalurl")))
	configuration.SetSpeedTestFileLocation(r.FormValue("stestfileloc"))
	configuration.SetPingDelay(r.FormValue("pingdelay"))
	configuration.SetSpeedTestDelay(r.FormValue("speedtestdelay"))

	configuration.WriteToFile()
}

// Loads the settings from the config file
func buildConfigStruct(cp *configPage) {
	cp.Port = configuration.GetPortNumber()
	cp.InternalAddr = configuration.GetInternalIPs()
	cp.ExternalAddr = configuration.GetExternalIPs()
	cp.ExternalURL = configuration.GetExternalURLs()
	cp.SpeedFileLoc = configuration.GetSpeedFileLocation()
	cp.ReportFileLoc = configuration.GetReportsLocation()
	cp.PingDelay = configuration.GetPingDelay()
	cp.SpeedDelay = configuration.GetSpeedDelay()
}

// Converts a textarea into a slice with one
// line per slice index
func lineToSlice(text string) []string {
	slice := strings.Split(text, "\n")

	return slice
}

// Converts the textarea text into a map of
// id=address for storing in the config struct
func ipToMap(text string) map[string]string {
	strmap := make(map[string]string)
	nlsplit := strings.Split(text, "\n")

	for _, v := range nlsplit {
		if v != "" {
			ls := strings.Split(v, "=")
			strmap[strings.ToLower(ls[0])] = ls[1]
		}
	}

	return strmap
}
