package reporter

import (
	"strconv"
)

type ReportData struct {
	Status          string
	RoomNumber      string
	Building        string
	RunLength       string
	StartTime       string // Date and Time
	Uptime          string
	LastConnect     string // Date and Time
	DisconnectCount int
	Timeline        string
	AvgSpeed        float32
	MaxSpeed        float32
	MinSpeed        float32
	Graph           string
}

func Fill(stringData map[string]string) ReportData {
	rd := ReportData{}

	rd.Status = "new"
	rd.RoomNumber = data["room"]
	rd.Building = data["building"]
	rd.RunLength = data["run-length"]
	rd.StartTime = data["start-time"]
	rd.Uptime = data["uptime"]
	rd.LastConnect = data["last-connection"]
	rd.DisconnectCount = Itoa(data["disconnections"])
	rd.Timeline = data["timeline"]
	rd.AvgSpeed, _ = ParseFloat(data["average-speed"], 32)
	rd.MaxSpeed, _ = ParseFloat(data["max-speed"], 32)
	rd.MinSpeed, _ = ParseFloat(data["min-speed"], 32)
	// rd.Graph = data["graph"]

	return rd
}
