package tools

import (
	"NetworkObserver/configuration"
	"NetworkObserver/reporter"
	"errors"
	"strconv"
	"strings"
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
	timeline := make([]string, 3)
	timeline[0] = ""
	timeline[1] = ""
	timeline[2] = ""

	lp := time.Now()
	end := time.Now().Add(time.Duration(runlen) * time.Hour)

	for time.Now().Before(end) {
		delay, _ := strconv.Atoi(pi.pingDelay)
		if time.Now().After(lp.Add(time.Duration(delay) * time.Second)) {
			pr := Ping(pi)
			lp = time.Now()

			if pr.internal {
				timeline[0] = timeline[0] + "o"
			} else {
				timeline[0] = timeline[0] + "*"
			}

			if pr.external_ip {
				timeline[1] = timeline[1] + "o"
			} else {
				timeline[1] = timeline[1] + "*"
			}

			if pr.external_url {
				timeline[2] = timeline[2] + "o"
			} else {
				timeline[2] = timeline[2] + "*"
			}

			// 3600 = number of seconds in an hour
			if updateReportTimeline(timeline, 10) { //3600/delay) {
				timeline[0] = ""
				timeline[1] = ""
				timeline[2] = ""
			}
		}
		// speedtest -- later
	}
}

func updateReportTimeline(tl []string, countPerHour int) bool {
	reset := false
	if countPerHour <= len(tl[0]) {
		if strings.Contains(tl[0], "*") {
			internalHourly += "*"
		} else {
			internalHourly += "o"
		}

		if strings.Contains(tl[1], "*") {
			externalIPHourly += "*"
		} else {
			externalIPHourly += "o"
		}

		if strings.Contains(tl[2], "*") {
			externalURLHourly += "*"
		} else {
			externalURLHourly += "o"
		}

		reset = true
	}

	tl_str := "Internal: " + tl[0] + "\nPast Hours: " + internalHourly +
		"External IP: " + tl[1] + "\nPast Hours: " + externalIPHourly +
		"External URL: " + tl[2] + "\nPast Hours: " + externalURLHourly

	reporter.SetTimeline(tl_str)

	return reset
}
