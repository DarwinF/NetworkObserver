package configuration

import (
	"errors"
	"math/rand"
	"strings"
)

func GetDeviceIP() string {
	return sysConfig.DeviceIP
}

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

func GetInternalIPbyKey(key string) (string, error) {
	for k, v := range sysConfig.InternalIPs {
		if strings.Contains(k, key) {
			return v, nil
		}
	}

	return "", errors.New("No ip associated with the key could be found.")
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
