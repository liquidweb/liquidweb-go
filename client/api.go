package client

import (
	network "git.liquidweb.com/masre/liquidweb-go/network"
	storm "git.liquidweb.com/masre/liquidweb-go/storm"
)

// API is the structure that houses all of our various API clients that interact with various Storm resources.
type API struct {
	NetworkDNS  network.DNSBackend
	NetworkZone network.ZoneBackend
	StormConfig storm.ConfigBackend
	StormServer storm.ServerBackend
	NetworkVIP  network.VIPBackend
}

// NewAPI is the API client for interacting with Storm.
func NewAPI(username string, password string, url string, timeout int) (*API, error) {
	config, err := NewConfig(username, password, url, timeout, true)
	if err != nil {
		return nil, err
	}

	// Initialize http backend
	client := NewClient(config)
	api := &API{
		NetworkDNS:  &network.DNSClient{Backend: client},
		NetworkZone: &network.ZoneClient{Backend: client},
		StormConfig: &storm.ConfigClient{Backend: client},
		StormServer: &storm.ServerClient{Backend: client},
		NetworkVIP:  &network.VIPClient{Backend: client},
	}

	return api, nil
}
