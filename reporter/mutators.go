package reporter

func SetUptime(ut string) {
	rd.Uptime = ut
}

func SetLastConnect(lc string) {
	rd.LastConnect = lc
}

func SetDisconnectCount(dc int) {
	rd.DisconnectCount = dc
}

func SetTimeline(tl string) {
	rd.Timeline = tl
}

func SetStatus(status string) {
	rd.Status = status
}

func SetLocation(loc string) {
	rd.Location = loc
}

func SetStartTime(st string) {
	rd.StartTime = st
}
