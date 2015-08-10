package tools

import (
	"NetworkObserver/configuration"
	"NetworkObserver/webserv"
)

type TestData struct {
	location        string
	runlen          string
	ext_ip          string
	ext_url         string
	speedtestfile   string
	ping_delay      string
	speedtest_delay string
}

func SetupTest(td TestData) pingInfo {
	pi := pingInfo{}

	if td.ext_ip == "" {
		pi.externalip = configuration.GetRandomExternalIP()
	} else {
		pi.externalip = td.ext_ip
	}

	if td.ext_url == "" {
		pi.externalurl = configuration.GetRandomExternalURL()
	} else {
		pi.externalurl = td.ext_url
	}

	pi.internalip = configuration.GetInternalIPbyKey(location)

	pi.pingDelay = td.ping_delay

	return pi
}

func RunTest(pi pingInfo, runlen int) {
	// set current time
	// start loop that lasts runlength long
	// ping
	// speedtest -- later
	// delay
}
