package configuration

func SetInternalIP(ip map[string]string) {
	sysConfig.InternalIPs = ip
}

func SetReportLocations(loc string) {
	sysConfig.ReportLocations = loc
}

func SetPortNumber(pn string) {
	sysConfig.PortNumber = pn
}

func SetExternalIPs(ips []string) {
	sysConfig.ExternalIP = ips
}

func SetExternalURLs(urls []string) {
	sysConfig.ExternalURL = urls
}

func SetSpeedTestFileLocation(loc string) {
	sysConfig.SpeedTestFileLocation = loc
}

func SetPingDelay(delay string) {
	sysConfig.PingDelay = delay
}

func SetSpeedTestDelay(delay string) {
	sysConfig.SpeedTestDelay = delay
}
