package configuration

import (
	"NetworkObserver/pkg/settings"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
var config Configuration

// read the .json file into the correct structs
func init() {
	cf := configPath
	file, err := ioutil.ReadFile("./config.json")

	// Todo: Grab the error from this
	json.Unmarshal(file, &config)
}

// WriteToFile - Writes the configuration to file
func WriteToFile() {
	if !updated {
		return
	}

	file, _ := os.Create(configPath)
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	data, err := json.Marshal(config)
	if err != nil {
		fmt.Println("error: ", err)
	}

	w.Write(data)

	// everything is written and nothing is new anymore
	updated = false
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

// SetPingDelay - Sets the ping delay
func SetPingDelay(delay string) {
	if delay != sysConfig.PingDelay {
		updated = true
		sysConfig.PingDelay = delay
	}
}

// SetSpeedTestDelay - Sets the speed tests delay
func SetSpeedTestDelay(delay string) {
	if delay != sysConfig.SpeedTestDelay {
		updated = true
		sysConfig.SpeedTestDelay = delay
	}
}

// GetDeviceIP - Returns the devices ip address
func GetDeviceIP() string {
	return sysConfig.DeviceIP
}

// GetPortNumber - Returns the port number
func GetPortNumber() string {
	return sysConfig.PortNumber
}

// GetInternalIPs returns all the internal ip addresses
func GetInternalIPs() string {
	ipString := ""

	for k, v := range sysConfig.InternalIPs {
		ipString += k + "=" + v + "\n"
	}

	return ipString
}

// GetInternalIPbyKey - Returns an internal ip address by searching based on the key passed in
func GetInternalIPbyKey(key string) (string, error) {
	for k, v := range sysConfig.InternalIPs {
		if strings.Contains(k, key) {
			return v, nil
		}
	}

	return "", errors.New("No ip associated with the key could be found.")
}

// GetExternalIPs - Returns all the external ip addresses
func GetExternalIPs() string {
	ipString := ""

	for _, v := range sysConfig.ExternalIP {
		ipString += v + "\n"
	}

	return ipString
}

// GetRandomExternalIP - Returns a random external IP address
func GetRandomExternalIP() string {
	return sysConfig.ExternalIP[rand.Intn(len(sysConfig.ExternalIP))]
}

// GetExternalURLs - Returns all external urls
func GetExternalURLs() string {
	urlString := ""

	for _, v := range sysConfig.ExternalURL {
		urlString += v + "\n"
	}

	return urlString
}

// GetRandomExternalURL - Returns a random external url
func GetRandomExternalURL() string {
	return sysConfig.ExternalURL[rand.Intn(len(sysConfig.ExternalURL))]
}

// GetSpeedFileLocation - Returns the location of the speed test files
func GetSpeedFileLocation() string {
	return sysConfig.SpeedTestFileLocation
}

// GetReportsLocation - Returns the location of the report files
func GetReportsLocation() string {
	return sysConfig.ReportLocations
}

// GetPingDelay - Returns the ping delay
func GetPingDelay() string {
	return sysConfig.PingDelay
}

// GetSpeedDelay - Returns the speed delay
func GetSpeedDelay() string {
	return sysConfig.SpeedTestDelay
}
