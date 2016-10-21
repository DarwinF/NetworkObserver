package configuration

import (
	logger "NetworkObserver/pkg/logging"
	"NetworkObserver/pkg/settings"
	"bufio"
	"encoding/json"
	"errors"
	"ioutil"
	"math/rand"
	"os"
	"reflect"
	"strings"
)

//--------------------------------
// Structs
//--------------------------------

// SystemSettings - General system settings
type SystemSettings struct {
	DeviceIP string
	Port     int
}

// DelaySettings - General test configuration settings
type DelaySettings struct {
	PingDelay      int
	SpeedTestDelay int
}

type taggedIP struct {
	ID      string
	Address string
}

type taggedURL struct {
	ID  string
	URL string
}

type taggedPath struct {
	ID   string
	Path string
}

// ExternalAddresses - External ip addresses and url addresses
type ExternalAddresses struct {
	IPAddresses  []taggedIP
	URLAddresses []taggedURL
}

// TestSettings - Settings for testing
type TestSettings struct {
	Configuration     DelaySettings
	InternalAddresses []taggedIP
	ExternalAddresses ExternalAddresses
	FileLocations     []taggedPath
}

// Configuration - all the configuration settings
type Configuration struct {
	SystemSettings SystemSettings
	TestSettings   TestSettings
}

//--------------------------------
// Variables
//--------------------------------
var loc = settings.AppLocation
var configPath = loc + "/config.json"
var sysConfig SystemSettings

var updated = false

// read the .ini file and fill the system struct
// with the data
func init() {
	cf := configPath

	// TODO: If the config file doesn't exist, generate a sample one

	file, err := os.Open(cf)

	if err != nil {
		logger.WriteString("The config file " + cf +
			" could not be found. A config file can be created by editing the configuration page.")
		return
	}

	defer file.Close()

	file, err := ioutil.ReadFile("./config.json")

	var config Configuration
	// Todo: Grab the error from this
	json.Unmarshal(file, &config)
}

// Write the configuration settings to the configuration file
func WriteToFile() {
	if !updated {
		return
	}

	file, _ := os.Create(configPath)
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	w.WriteString("[General]\n")
	w.WriteString("deviceip=" + sysConfig.DeviceIP + "\n")
	w.WriteString("port=" + sysConfig.PortNumber + "\n")

	w.WriteString("[Internal Addresses]\n")
	writeMap(sysConfig.InternalIPs, w)

	w.WriteString("[External Addresses]\n")
	writeSlice(sysConfig.ExternalIP, "ip", w)
	writeSlice(sysConfig.ExternalURL, "url", w)

	w.WriteString("[File Locations]\n")
	w.WriteString("speedtest=" + sysConfig.SpeedTestFileLocation + "\n")
	w.WriteString("reports=" + sysConfig.ReportLocations + "\n")

	w.WriteString("[Test Delays]\n")
	w.WriteString("ping=" + sysConfig.PingDelay + "\n")
	w.WriteString("speedtest=" + sysConfig.SpeedTestDelay + "\n")

	// everything is written and nothing is new anymore
	updated = false
}

func writeMap(m map[string]string, w *bufio.Writer) {
	for k, v := range m {
		w.WriteString(k + "=" + v + "\n")
	}
}

func writeSlice(s []string, id string, w *bufio.Writer) {
	for _, v := range s {
		w.WriteString(id + "=" + v + "\n")
	}
}

func identify(line string, currSect Section) Section {
	if line[0] == '[' {
		str := strings.Replace(line, "[", "", -1)
		str = strings.Replace(str, "]", "", -1)

		switch str {
		case "Internal Addresses":
			return InternalAddr
		case "External Addresses":
			return ExternalAddr
		case "General":
			return General
		case "File Locations":
			return FileLocs
		case "Test Delays":
			return TestDelays
		}
	}

	return currSect
}

func storeValue(line string, sect Section) {
	str := make([]string, 0)
	// Seprate the line into the two parts: identifier and value
	// if the line is a header line, ignore it
	if line[0] != '[' && line[0] != ';' {
		str = strings.Split(line, "=")
	} else {
		return
	}
	// store the value in the correct struct element
	if sect == InternalAddr {
		sysConfig.InternalIPs[strings.ToLower(str[0])] = str[1]
	} else if sect == ExternalAddr {
		if str[0] == "ip" {
			sysConfig.ExternalIP = append(sysConfig.ExternalIP, str[1])
		} else if str[0] == "url" {
			sysConfig.ExternalURL = append(sysConfig.ExternalURL, str[1])
		}
	} else if sect == General {
		if str[0] == "port" {
			sysConfig.PortNumber = str[1]
		} else if str[0] == "deviceip" {
			sysConfig.DeviceIP = str[1]
		}
	} else if sect == FileLocs {
		if str[0] == "speedtest" {
			sysConfig.SpeedTestFileLocation = str[1]
		} else if str[0] == "reports" {
			sysConfig.ReportLocations = str[1]
		}
	} else if sect == TestDelays {
		if str[0] == "ping" {
			sysConfig.PingDelay = str[1]
		} else if str[0] == "speedtest" {
			sysConfig.SpeedTestDelay = str[1]
		}
	}
}

func SetInternalIP(ip map[string]string) {
	equal := reflect.DeepEqual(ip, sysConfig.InternalIPs)

	if !equal {
		updated = true
		sysConfig.InternalIPs = ip
	}
}

func SetReportLocations(loc string) {
	if loc != sysConfig.ReportLocations {
		updated = true
		sysConfig.ReportLocations = loc
	}
}

func SetDeviceIP(ip string) {
	if ip != sysConfig.DeviceIP {
		updated = true
		sysConfig.DeviceIP = ip
	}
}

func SetPortNumber(pn string) {
	if pn != sysConfig.PortNumber {
		updated = true
		sysConfig.PortNumber = pn
	}
}

func SetExternalIPs(ips []string) {
	equal := reflect.DeepEqual(ips, sysConfig.ExternalIP)

	if !equal {
		updated = true
		sysConfig.ExternalIP = ips
	}
}

func SetExternalURLs(urls []string) {
	equal := reflect.DeepEqual(urls, sysConfig.ExternalURL)

	if !equal {
		updated = true
		sysConfig.ExternalURL = urls
	}
}

func SetSpeedTestFileLocation(loc string) {
	if loc != sysConfig.SpeedTestFileLocation {
		updated = true
		sysConfig.SpeedTestFileLocation = loc
	}
}

func SetPingDelay(delay string) {
	if delay != sysConfig.PingDelay {
		updated = true
		sysConfig.PingDelay = delay
	}
}

func SetSpeedTestDelay(delay string) {
	if delay != sysConfig.SpeedTestDelay {
		updated = true
		sysConfig.SpeedTestDelay = delay
	}
}

func GetDeviceIP() string {
	return sysConfig.DeviceIP
}

func GetPortNumber() string {
	return sysConfig.PortNumber
}

func GetInternalIPs() string {
	ipString := ""

	for k, v := range sysConfig.InternalIPs {
		ipString += k + "=" + v + "\n"
	}

	return ipString
}

func GetInternalIPbyKey(key string) (string, error) {
	for k, v := range sysConfig.InternalIPs {
		if strings.Contains(k, key) {
			return v, nil
		}
	}

	return "", errors.New("No ip associated with the key could be found.")
}

func GetExternalIPs() string {
	ipString := ""

	for _, v := range sysConfig.ExternalIP {
		ipString += v + "\n"
	}

	return ipString
}

func GetRandomExternalIP() string {
	return sysConfig.ExternalIP[rand.Intn(len(sysConfig.ExternalIP))]
}

func GetExternalURLs() string {
	urlString := ""

	for _, v := range sysConfig.ExternalURL {
		urlString += v + "\n"
	}

	return urlString
}

func GetRandomExternalURL() string {
	return sysConfig.ExternalURL[rand.Intn(len(sysConfig.ExternalURL))]
}

func GetSpeedFileLocation() string {
	return sysConfig.SpeedTestFileLocation
}

func GetReportsLocation() string {
	return sysConfig.ReportLocations
}

func GetPingDelay() string {
	return sysConfig.PingDelay
}

func GetSpeedDelay() string {
	return sysConfig.SpeedTestDelay
}
