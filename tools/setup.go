package tools

import (
	"NetworkObserver/configuration"
	"errors"
	"strconv"
	"time"
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

type pingResponse struct {
	internal     bool
	external_ip  bool
	external_url bool

	err error
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
	lp := time.Now()
	end := time.Now().Add(time.Duration(runlen) * time.Hour)

	for time.Now().Before(end) {
		delay, _ := strconv.Atoi(pi.pingDelay)
		if time.Now().After(lp.Add(time.Duration(delay) * time.Second)) {
			Ping(pi)
			lp = time.Now()

			// Log response in report xml file
		}
		// speedtest -- later
	}
}
