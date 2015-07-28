package configuration

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//--------------------------------
// Structs
//--------------------------------

// this struct contains settings
// that are configured for a specific
// test run
type TestSettings struct {
	InternalIP string
	RunLength  string

	configSettings
}

// these are settings that are global
// but can also be overwritten for one
// specific test run
type configSettings struct {
	ExternalIP            []string
	ExternalURL           []string
	SpeedTestFileLocation string
	PingDelay             string
	SpeedTestDelay        string
	PortNumber            string
}

// these are the settings that are globally
// set for the program
type SystemSettings struct {
	InternalIPs     map[string]string
	ReportLocations string

	configSettings
}

//--------------------------------
// Variables
//--------------------------------
var configPath = "sample_config.txt"
var sysConfig SystemSettings

//--------------------------------
// Enum
//--------------------------------
type Section int8

const (
	Pretext      Section = -1
	General      Section = 1
	InternalAddr Section = 2
	ExternalAddr Section = 3
	FileLocs     Section = 4
	TestDelays   Section = 5
)

// read the .ini file and fill the system struct
// with the data
func init() {
	// Setup struct
	sysConfig = SystemSettings{}
	sysConfig.InternalIPs = make(map[string]string)
	sysConfig.ExternalIP = make([]string, 0)
	sysConfig.ExternalURL = make([]string, 0)

	file, _ := os.Open(configPath)
	// report error if there is one
	defer file.Close()

	// Set the initial value for sect, this way we have a base
	// value
	sect := Pretext

	// read the line from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		// identify what section we are in
		if str != "" {
			sect = identify(str, sect)
			storeValue(str, sect)
		}
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
		sysConfig.InternalIPs[str[0]] = str[1]
	} else if sect == ExternalAddr {
		if str[0] == "ip" {
			sysConfig.ExternalIP = append(sysConfig.ExternalIP, str[1])
		} else if str[0] == "url" {
			sysConfig.ExternalURL = append(sysConfig.ExternalURL, str[1])
		}
	} else if sect == General {
		if str[0] == "port" {
			sysConfig.PortNumber = str[1]
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
