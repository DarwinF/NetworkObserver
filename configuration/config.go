package configuration

import ()

type ConfigData struct {
	InternalIP_UCR string
	InternalIP_MSH string
	InternalIP_AVM string
	InternalIP_PHR string
	ExternalDomain string
	ExternalIP     string
	PingDelay      string // time or int?
	SpeedTestDelay string
}

// For simple interaction with XML
func (cd *ConfigData) CreateMap() map[string]string {
	m := make(map[string]string)

	return m
}
