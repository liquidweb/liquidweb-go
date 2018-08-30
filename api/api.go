package api

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	networkzone "git.liquidweb.com/masre/liquidweb-go/networkzone"
	stormconfig "git.liquidweb.com/masre/liquidweb-go/storm/config"
	stormserver "git.liquidweb.com/masre/liquidweb-go/storm/server"
)

// API is the structure that houses all of our various API clients that interact with various Storm resources.
type API struct {
	NetworkZone *networkzone.Client
	StormConfig *stormconfig.Client
	StormServer *stormserver.Client
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
		NetworkZone: &networkzone.Client{Backend: client},
		StormConfig: &stormconfig.Client{Backend: client},
		StormServer: &stormserver.Client{Backend: client},
	}

	return api, nil
}
