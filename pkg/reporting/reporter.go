package reporter

import (
	"os"

	"github.com/darwinfroese/networkobserver/pkg/logging"
	"github.com/darwinfroese/networkobserver/pkg/settings"
)

type connection struct {
	Uptime          string
	LastConnect     string
	DisconnectCount int
	Timeline        string
}

type speedTest struct {
	AvgSpeed float32
	MaxSpeed float32
	MinSpeed float32
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
