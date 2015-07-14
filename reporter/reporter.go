package reporter

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

var currentReport *ReportData = nil

// Read data from report file and store it into a ReportData sturcture
func CreateNew(roomNumber, building string, rd *ReportData) {
	rd.Status = "Run"
	rd.RoomNumber = roomNumber
	rd.Building = building
	rd.RunLength = "1000hours"
	rd.StartTime = "Yesterday i think"
	rd.Uptime = "Ummm... maybe ten minutes"
	rd.LastConnect = "Not that long ago actually"
	rd.DisconnectCount = 100
	rd.Timeline = "A little bit of this a little bit of that"
	rd.AvgSpeed = 100.0
	rd.MaxSpeed = 10000.0
	rd.MinSpeed = 1.0
	rd.Graph = "100"

	currentReport = rd
}
