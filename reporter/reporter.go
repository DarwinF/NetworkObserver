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

}

// For simple interaction with XML
func (rd *ReportData) CreateMap() map[string]string {
	return nil
}
