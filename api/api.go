package api

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	networkzone "git.liquidweb.com/masre/liquidweb-go/networkzone"
	stormserver "git.liquidweb.com/masre/liquidweb-go/storm/server"
)

// API is the structure that houses all of our various API clients that interact with various Storm resources.
type API struct {
	NetworkZone networkzone.NetworkZoneClient
	StormServer stormserver.StormServerClient
}

// NewAPI is the API client for interacting with Storm.
func NewAPI(username string, password string, url string, timeout int) (*API, error) {
	config, err := client.NewConfig(username, password, url, timeout, true)
	if err != nil {
		return nil, err
	}

	// Initialize http backend
	client := client.NewClient(config)
	api := &API{
		NetworkZone: &networkzone.Client{Client: client},
		StormServer: &stormserver.Client{Client: client},
	}

	return api, nil
}
