package reporter

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
