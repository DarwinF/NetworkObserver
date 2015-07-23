package reporter

/*
import (
	"strconv"
)*/

type Location struct {
	RoomNumber string `xml:"Location>RoomNumber"`
	Building   string `xml:"Location>Building"`
}

type Connection struct {
	Uptime          string `xml:"Connection>Uptime"`
	LastConnect     string `xml:"Connection>LastConnect"`
	DisconnectCount int    `xml:"Connection>DisconnectCount"`
	Timeline        string `xml:"Connection>Timeline"`
}

type SpeedTest struct {
	AvgSpeed float32 `xml:"SpeedTest>AvgSpeed"`
	MaxSpeed float32 `xml:"SpeedTest>MaxSpeed"`
	MinSpeed float32 `xml:"SpeedTest>MinSpeed"`
}

type ReportData struct {
	Status string
	Location
	StartTime string
	Connection
	SpeedTest
	Graph string
}

/*
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
*/
