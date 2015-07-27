package configuration

// this struct contains settings
// that are configured for a specific
// test run
type TestSettings struct {
	InternalIP string
	RunLength  string

	configSettings
}

// these are settings that are global
// but can also be overwritten for one
// specific test run
type configSettings struct {
	ExternalIP            []string
	ExternalURL           []string
	SpeedTestFileLocation string
	PingDelay             string
	SpeedTestDelay        string
	PortNumber            string
}

// these are the settings that are globally
// set for the program
type SystemSettings struct {
	InternalIPs     map[string]string
	ReportLocations string

	configSettings
}

// read the .ini file and fill the system struct
// with the data
func init() {

}
