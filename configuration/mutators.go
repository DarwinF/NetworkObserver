package configuration

import (
	"reflect"
)

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
