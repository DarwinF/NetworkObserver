package reporter

import (
	"NetworkObserver/logger"
	"os"
)

type connection struct {
	Uptime          string `xml:"Connection>Uptime"`
	LastConnect     string `xml:"Connection>LastConnect"`
	DisconnectCount int    `xml:"Connection>DisconnectCount"`
	Timeline        string `xml:"Connection>Timeline"`
}

type speedTest struct {
	AvgSpeed float32 `xml:"SpeedTest>AvgSpeed"`
	MaxSpeed float32 `xml:"SpeedTest>MaxSpeed"`
	MinSpeed float32 `xml:"SpeedTest>MinSpeed"`
}

type ReportData struct {
	Status    string
	Location  string
	StartTime string
	Graph     string
	connection
	speedTest
}

var file *os.File
var rd ReportData

func init() {
	var err error
	file, err = os.OpenFile("report.xml", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		logger.WriteString("There was a problem opening the report.xml file")
	}

	rd = ReportData{}
}
