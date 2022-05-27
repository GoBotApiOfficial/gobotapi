package gobotapi

import "fmt"

type Config struct {
	HostName string
	Port     int
	Https    bool
}

func (config Config) link() string {
	var protocol string
	if len(config.HostName) == 0 {
		config.HostName = "api.telegram.org"
		config.Https = true
	}
	if config.Https {
		protocol = "https"
	} else {
		protocol = "http"
	}
	if config.Port == 0 {
		if config.Https {
			config.Port = 443
		} else {
			config.Port = 80
		}
	}
	return fmt.Sprintf("%s://%s:%d/", protocol, config.HostName, config.Port)
}
