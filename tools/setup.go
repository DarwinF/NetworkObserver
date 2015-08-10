package tools

import (
	"NetworkObserver/configuration"
	"errors"
)

type TestData struct {
	Location        string
	Runlen          string
	Ext_ip          string
	Ext_url         string
	Speedtestfile   string
	Ping_delay      string
	Speedtest_delay string
}

func SetupTest(td TestData) (pingInfo, error) {
	pi := pingInfo{}
	err := errors.New("")

	if td.Ext_ip == "" {
		pi.externalip = configuration.GetRandomExternalIP()
	} else {
		pi.externalip = td.Ext_ip
	}

	if td.Ext_url == "" {
		pi.externalurl = configuration.GetRandomExternalURL()
	} else {
		pi.externalurl = td.Ext_url
	}

	pi.internalip, err = configuration.GetInternalIPbyKey(td.Location)
	if err != nil {
		return pingInfo{}, err
	}

	pi.pingDelay = td.Ping_delay

	return pi, nil
}

func RunTest(pi pingInfo, runlen int) {
	// set current time
	// start loop that lasts runlength long
	// ping
	// speedtest -- later
	// delay
}
