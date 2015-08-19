package reporter

import (
	"NetworkObserver/logger"
	"NetworkObserver/settings"
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
var loc string = settings.AppLocation + "/Reports/"
var rf string = loc + "report.xml"

func init() {
	if _, err := os.Stat(loc); os.IsNotExist(err) {
		logger.WriteString("The reports folder does not exist. Creating the folder now.")
		os.Mkdir(loc, 0777)
	}

	var e error
	file, e = os.OpenFile(rf, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if e != nil {
		logger.WriteString("There was a problem opening the report.xml file")
	}

	rd = ReportData{}
}
