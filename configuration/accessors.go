package configuration

import (
	"math/rand"
)

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

func GetInternalIPbyKey(k string) string {
	v = ""

	if k != "" {
		v = sysConfig.InternalIPs[k]
	}

	if v == "" {
		return sysConfig.InternalIPs[rand.Intn(len(sysConfig.InternalIPs))]
	} else {
		return v
	}
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
