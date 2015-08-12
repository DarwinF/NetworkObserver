package reporter

func GetUptime() string {
	return rd.Uptime
}

func GetLastConnect() string {
	return rd.LastConnect
}

func GetDisconnectCount() int {
	return rd.DisconnectCount
}

func GetTimeline() string {
	return rd.Timeline
}

func GetStatus() string {
	return rd.Status
}

func GetLocation() string {
	return rd.Location
}

func GetStartTime() string {
	return rd.StartTime
}
