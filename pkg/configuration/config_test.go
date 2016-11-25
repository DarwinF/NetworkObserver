package configuration

import (
	"os"
	"reflect"
	"testing"
)

var configPath = "./config.json"
var configurator Configurator

var defaultConfiguration = Configuration{
	SystemSettings: SystemSettings{
		DeviceIP: "127.0.0.1",
		Port:     "8080",
	},
	TestSettings: TestSettings{
		Configuration: DelaySettings{
			PingDelay:      "60",
			SpeedTestDelay: "60",
		},
		InternalAddresses: []taggedIP{
			taggedIP{ID: "Localhost", Address: "127.0.0.1"},
		},
		ExternalAddresses: ExternalAddresses{
			IPAddresses: []taggedIP{
				taggedIP{ID: "Google DNS", Address: "8.8.8.8"},
			},
			URLAddresses: []taggedURL{
				taggedURL{ID: "Google", URL: "www.google.ca"},
			},
		},
		FileLocations: []taggedPath{
			taggedPath{ID: "Speed Test Files", Path: "./Reports/SpeedTests/"},
			taggedPath{ID: "Ping Test Files", Path: "./Reports/PingTests/"},
		},
	},
}

func Test_ConfigFileMissingGeneratesDefaultConfig(t *testing.T) {
	setupConfigurator()

	t.Log("\nGot: ", configurator.GetSettings(), "\nExp: ", defaultConfiguration)
	removeConfigFile()

	if !reflect.DeepEqual(configurator.GetSettings(), defaultConfiguration) {
		t.Fatal("The generated configuration doesn't match the default configuration")
	}
}

func setupConfigurator() {
	configurator = NewConfigurator()
}

func removeConfigFile() {
	if _, err := os.Stat(defaultConfigPath); os.IsNotExist(err) {
		return
	}

	os.Remove(defaultConfigPath)
}
