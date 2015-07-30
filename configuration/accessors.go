package configuration

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

func GetExternalIPs() string {
	ipString := ""

	for _, v := range sysConfig.ExternalIP {
		ipString += v + "\n"
	}

	return ipString
}

func GetExternalURLs() string {
	urlString := ""

	for _, v := range sysConfig.ExternalURL {
		urlString += v + "\n"
	}

	return urlString
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
