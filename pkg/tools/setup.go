package tools

import (
	"NetworkObserver/pkg/configuration"
	logger "NetworkObserver/pkg/logging"
	"NetworkObserver/pkg/reporting"
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

var internalHourly = ""
var externalIPHourly = ""
var externalURLHourly = ""

func SetupTest(td TestData) (pingInfo, error) {
	internalHourly = ""
	externalIPHourly = ""
	externalURLHourly = ""

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
	tl := map[string]int{
		"internal": 0,
		"extip":    0,
		"exturl":   0,
		"total":    0,
	}

	lp := time.Now()
	end := time.Now().Add(time.Duration(runlen) * time.Hour)

	logger.WriteString("Starting the test...")

	for time.Now().Before(end) {
		delay, _ := strconv.Atoi(pi.pingDelay)
		if time.Now().After(lp.Add(time.Duration(delay) * time.Second)) {
			pr := Ping(pi)
			lp = time.Now()

			if pr.internal {
				tl["internal"]++
			}
			if pr.external_ip {
				tl["extip"]++
			}
			// if pr.external_url {
			// 	tl["exturl"]++
			// }
			tl["total"]++

			timeline := "Internal Ping Success Count: " + strconv.Itoa(tl["internal"]) + "/" + strconv.Itoa(tl["total"]) +
				"\nExternal IP Ping Success Count: " + strconv.Itoa(tl["extip"]) + "/" + strconv.Itoa(tl["total"]) //+
			// "\nExternal URL Ping Success Count: " + strconv.Itoa(tl["exturl"]) + "/" + strconv.Itoa(tl["total"])

			reporter.SetTimeline(timeline)
		}
		// speedtest -- later
	}

	logger.WriteString("The test has been completed.")
}
