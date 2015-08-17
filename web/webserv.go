//--------------------------------------------
// web/webserv.go
//
// Handles serving and authenticating for all
// the webpages.
//
// All the handler functions are declared in
// this file.
//--------------------------------------------

package web

import (
	"NetworkObserver/auth"
	"NetworkObserver/configuration"
	"NetworkObserver/logger"
	"NetworkObserver/reporter"
	"NetworkObserver/tools"
	"crypto/sha256"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//--------------------------------------------
// Variables
//--------------------------------------------
var errMsg string = ""
var loc string = "/var/lib/apps/NetworkObserver/HTML/"

//--------------------------------------------
// Structs for Pages
//--------------------------------------------
type configPage struct {
	DeviceIP      string
	Port          string
	InternalAddr  string
	ExternalAddr  string
	ExternalURL   string
	SpeedFileLoc  string
	ReportFileLoc string
	PingDelay     string
	SpeedDelay    string
}

type testPage struct {
	ErrorMessage
	SpeedTestFileLocation string
	PingDelay             string
	SpeedTestDelay        string
}

type createAccount struct {
	ErrorMessage
	Username string
}

type ErrorMessage struct {
	Msg string
}

// All URLs default to this function
func Root(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		servePageStatic(w, r, loc+"login.html")
	}
}

// Handles URLs referencing dashboard
func Dashboard(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid {
		servePageStatic(w, r, loc+"dashboard.html")
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
	// Hash password
	hash := sha256.Sum256([]byte(pword))
	authenticated := auth.CheckCredentials(uname, hash)

	if authenticated == true {
		logger.WriteString("User \"" + uname + "\" authenticated successfully.")
		auth.SetSessionID(w)
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		logger.WriteString("User \"" + uname + "\" did not authenticate successfully.")
		servePageStatic(w, r, loc+"error.html")
	}
}

// Create a new account by comparing both of the passwords
// entered and then hashing and storing the password
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	servePageStatic(w, r, loc+"createaccount.html")
}

// Serve the webpage for creating an account
func HandleAccount(w http.ResponseWriter, r *http.Request) {
	un := r.FormValue("username")
	pw := r.FormValue("password")
	pwv := r.FormValue("password-verify")

	if auth.UsernameInUse(un) {
		ca := createAccount{}
		ca.Msg = "The username " + un + " is currently in use."
		ca.Username = ""
		servePageDynamic(w, r, loc+"createaccount.html", ca)
	} else if un == "" {
		ca := createAccount{}
		ca.Msg = "You must enter a username."
		ca.Username = ""
		servePageDynamic(w, r, loc+"createaccount.html", ca)
	} else if pw != pwv || pw == "" || pwv == "" {
		ca := createAccount{}
		ca.Msg = "The passwords do not match."
		ca.Username = un
		servePageDynamic(w, r, loc+"createaccount.html", ca)
	} else if pw == pwv {
		h := sha256.Sum256([]byte(pw))
		auth.SavePassword(un, h)
		servePageStatic(w, r, loc+"success.html")
	}
}

// Save the configuration settings and then reload the configuration
// page (which will automatically reload the new settings)
func SaveConfig(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid {
		// Save Configuration Settings
		saveConfigToStruct(r)
		http.Redirect(w, r, "/dashboard/configure", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// The test has had it's "test specific settings "
func SaveTest(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid {
		if r.FormValue("runlength") == "" || r.FormValue("location") == "" {
			errMsg = "You need to enter a run length and a location!"
			http.Redirect(w, r, "/dashboard/start_test", http.StatusFound)
		} else {
			td := tools.TestData{}
			td.Runlen = r.FormValue("runlength")
			td.Ext_ip = r.FormValue("externalIP")
			td.Ext_url = r.FormValue("externalURL")
			td.Location = r.FormValue("location")
			td.Ping_delay = r.FormValue("pingdelay")
			td.Speedtest_delay = r.FormValue("speedtestedelay")
			td.Speedtestfile = r.FormValue("stestfileloc")

			pi, err := tools.SetupTest(td)

			if err != nil {
				errMsg = "No ip associated with the key could be found."
				http.Redirect(w, r, "/dashboard/start_test", http.StatusFound)
			}

			reporter.SetStartTime(time.Now().Format("Jan 02, 2006 - 15:04"))
			reporter.SetLocation(td.Location)

			runlen, _ := strconv.Atoi(td.Runlen)
			go tools.RunTest(pi, runlen)
			http.Redirect(w, r, "/dashboard/reports", http.StatusFound)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// A redirection page to notify the user that the test has been started
func TestStarted(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid {
		servePageStatic(w, r, loc+"dashboard/testnotify.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	auth.RemoveCookie(w, r)

	http.Redirect(w, r, "/", http.StatusFound)
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

	if valid {
		servePageStatic(w, r, loc+"dashboard/settings.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Configure(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	// Build struct to serve page with:
	configStruct := configPage{}
	buildConfigStruct(&configStruct)

	if valid {
		servePageDynamic(w, r, loc+"dashboard/config.html", configStruct)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func StartTest(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	ts := testPage{}
	buildTestStruct(&ts)

	// Copy the error message over and display it and then erase it
	// this way the message will only be displayed when there is an
	// actual error
	if errMsg != "" {
		ts.Msg = errMsg
		errMsg = ""
	}

	if valid {
		servePageDynamic(w, r, loc+"dashboard/starttest.html", ts)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Reports(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid {
		rd := reporter.ReportData{}
		buildReport(&rd)
		servePageDynamic(w, r, loc+"dashboard/reports.html", rd)
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

// Saves the settings to the config file
func saveConfigToStruct(r *http.Request) {
	configuration.SetDeviceIP(r.FormValue("deviceip"))
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
	cp.DeviceIP = configuration.GetDeviceIP()
	cp.Port = configuration.GetPortNumber()
	cp.InternalAddr = configuration.GetInternalIPs()
	cp.ExternalAddr = configuration.GetExternalIPs()
	cp.ExternalURL = configuration.GetExternalURLs()
	cp.SpeedFileLoc = configuration.GetSpeedFileLocation()
	cp.ReportFileLoc = configuration.GetReportsLocation()
	cp.PingDelay = configuration.GetPingDelay()
	cp.SpeedDelay = configuration.GetSpeedDelay()
}

func buildTestStruct(ts *testPage) {
	ts.PingDelay = configuration.GetPingDelay()
	ts.SpeedTestDelay = configuration.GetSpeedDelay()
	ts.SpeedTestFileLocation = configuration.GetSpeedFileLocation()
}

func buildReport(r *reporter.ReportData) {
	r.Uptime = reporter.GetUptime()
	r.LastConnect = reporter.GetLastConnect()
	r.DisconnectCount = reporter.GetDisconnectCount()
	r.Timeline = reporter.GetTimeline()
	r.Status = reporter.GetStatus()
	r.Location = reporter.GetLocation()
	r.StartTime = reporter.GetStartTime()
}

// Converts a textarea into a slice with one
// line per slice index
func lineToSlice(text string) []string {
	// Newlines are considered "whitespace" in Go so this will remove the trailing
	// newline characters. Any number of blank lines can be inserted at the end of
	// the text area and they will be removed
	text = strings.TrimSpace(text)
	slice := strings.Split(text, "\n")

	return slice
}

// Converts the textarea text into a map of
// id=address for storing in the config struct
func ipToMap(text string) map[string]string {
	text = strings.TrimSpace(text)

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
