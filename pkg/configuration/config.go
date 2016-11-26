package configuration

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

//--------------------------------
// Structs
//--------------------------------

// Configurator interface
type Configurator interface {
	Update(*Configuration) (bool, error)
	GetSettings() Configuration
}

type configAdapter struct {
	settings Configuration
}

// SystemSettings - General system settings
type SystemSettings struct {
	DeviceIP string
	Port     string
}

// DelaySettings - General test configuration settings
type DelaySettings struct {
	PingDelay      string
	SpeedTestDelay string
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
var defaultConfigPath = "./config.json"

var updated = false
var config Configuration

// read the .json file into the correct structs
func setup() {
	updated = true

	if _, err := os.Stat(defaultConfigPath); os.IsNotExist(err) {
		config = generateDefaultConfig()
		writeToFile()
		return
	}

	file, _ := ioutil.ReadFile(defaultConfigPath)
	err := json.Unmarshal(file, &config)

	if err != nil {
		fmt.Printf("There was an error reading the config.json file located at: %s\n", defaultConfigPath)
		config = generateDefaultConfig()
		writeToFile()
	}
}

// NewConfigurator creates a new configurator with the default or
// saved configuration options
func NewConfigurator() Configurator {
	ca := configAdapter{}

	setup()
	ca.settings = config

	return &ca
}

// Update is used to update the current configuration
func (c *configAdapter) Update(config *Configuration) (bool, error) {

	return false, nil
}

func (c *configAdapter) GetSettings() Configuration {
	return c.settings
}

func generateDefaultConfig() (c Configuration) {
	c = Configuration{
		SystemSettings: SystemSettings{
			DeviceIP: "127.0.0.1",
			Port:     "8080",
		},
		TestSettings: TestSettings{
			Configuration: DelaySettings{
				PingDelay:      "60",
				SpeedTestDelay: "60",
			},
			InternalAddresses: []taggedIP{
				taggedIP{ID: "Localhost", Address: "127.0.0.1"},
			},
			ExternalAddresses: ExternalAddresses{
				IPAddresses: []taggedIP{
					taggedIP{ID: "Google DNS", Address: "8.8.8.8"},
				},
				URLAddresses: []taggedURL{
					taggedURL{ID: "Google", URL: "www.google.ca"},
				},
			},
			FileLocations: []taggedPath{
				taggedPath{ID: "Speed Test Files", Path: "./Reports/SpeedTests/"},
				taggedPath{ID: "Ping Test Files", Path: "./Reports/PingTests/"},
			},
		},
	}

	return
}

func writeToFile() {
	if !updated {
		return
	}

	file, err := os.Create(defaultConfigPath)
	if err != nil {
		fmt.Println("There was an error creating the file: ", err)
	}
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

// SetInternalIP - Sets the internal ip value
func SetInternalIP(ip map[string]string) {
	equal := reflect.DeepEqual(ip, config.TestSettings.InternalAddresses)

	if !equal {
		updated = true
		var ips []taggedIP
		//Todo: Write a mapping function for this
		for k, v := range ip {
			ips = append(ips, taggedIP{ID: k, Address: v})
		}

		config.TestSettings.InternalAddresses = ips
	}
}

// SetReportLocations - Sets the report location
func SetReportLocations(loc string) {

	// Todo: fix this -- and check for null list
	if loc != config.TestSettings.FileLocations[0].Path {
		updated = true
		config.TestSettings.FileLocations[0].Path = loc
	}
}

// SetDeviceIP - Sets the device IP
func SetDeviceIP(ip string) {
	if ip != config.SystemSettings.DeviceIP {
		updated = true
		config.SystemSettings.DeviceIP = ip
	}
}

// SetPortNumber - Sets the applications port number
func SetPortNumber(pn string) {
	// Todo: Cast to int
	if pn != config.SystemSettings.Port {
		updated = true
		config.SystemSettings.Port = pn
	}
}

// SetExternalIPs - Sets the external IP addresses
func SetExternalIPs(ips []string) {
	// Todo: This ended up getting mapped weird
}

// SetExternalURLs - sets the external URLS
func SetExternalURLs(urls []string) {
	// Todo: This also ended up getting mapped weird
}

// SetSpeedTestFileLocation - Sets the speed test file location
func SetSpeedTestFileLocation(loc string) {
	// Todo: fix this -- and check for null list
	if loc != config.TestSettings.FileLocations[1].Path {
		updated = true
		config.TestSettings.FileLocations[1].Path = loc
	}
}

// SetPingDelay - Sets the ping delay
func SetPingDelay(delay string) {
	// Todo: Cast to int
	if delay != config.TestSettings.Configuration.PingDelay {
		updated = true
		config.TestSettings.Configuration.PingDelay = delay
	}
}

// SetSpeedTestDelay - Sets the speed tests delay
func SetSpeedTestDelay(delay string) {
	if delay != config.TestSettings.Configuration.SpeedTestDelay {
		updated = true
		config.TestSettings.Configuration.SpeedTestDelay = delay
	}
}

// GetDeviceIP - Returns the devices ip address
func GetDeviceIP() string {
	return config.SystemSettings.DeviceIP
}

// GetPortNumber - Returns the port number
func GetPortNumber() string {
	return config.SystemSettings.Port
}

// GetInternalIPs - returns all the internal ip addresses
func GetInternalIPs() string {
	// Todo: Return this properly
	return ""
	// ipString := ""

	// for k, v := range config.TestSettings.InternalAddresses {
	// 	ipString += k + "=" + v + "\n"
	// }

	// return ipString
}

// GetInternalIPbyKey - Returns an internal ip address by searching based on the key passed in
func GetInternalIPbyKey(key string) (string, error) {
	// Todo: return this properly
	// for k, v := range sysConfig.InternalIPs {
	// 	if strings.Contains(k, key) {
	// 		return v, nil
	// 	}
	// }

	return "", errors.New("No ip associated with the key could be found.")
}

// GetExternalIPs - Returns all the external ip addresses
func GetExternalIPs() string {
	// Todo: return this properly
	// ipString := ""

	// for _, v := range sysConfig.ExternalIP {
	// 	ipString += v + "\n"
	// }

	// return ipString
	return ""
}

// GetRandomExternalIP - Returns a random external IP address
func GetRandomExternalIP() string {
	// Todo: return this properly
	return ""
	// return sysConfig.ExternalIP[rand.Intn(len(sysConfig.ExternalIP))]
}

// GetExternalURLs - Returns all external urls
func GetExternalURLs() string {
	urlString := ""

	// Todo: return this properly
	// for _, v := range sysConfig.ExternalURL {
	// 	urlString += v + "\n"
	// }

	return urlString
}

// GetRandomExternalURL - Returns a random external url
func GetRandomExternalURL() string {
	// Todo: return this properly
	return ""
	//return sysConfig.ExternalURL[rand.Intn(len(sysConfig.ExternalURL))]
}

// GetSpeedFileLocation - Returns the location of the speed test files
func GetSpeedFileLocation() string {
	return config.TestSettings.FileLocations[1].Path
}

// GetReportsLocation - Returns the location of the report files
func GetReportsLocation() string {
	return config.TestSettings.FileLocations[0].Path
}

// GetPingDelay - Returns the ping delay
func GetPingDelay() string {
	return config.TestSettings.Configuration.PingDelay
}

// GetSpeedDelay - Returns the speed delay
func GetSpeedDelay() string {
	return config.TestSettings.Configuration.SpeedTestDelay
}
