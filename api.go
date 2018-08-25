package storm

import (
	stormserver "git.liquidweb.com/cnichols/storm-client/storm/server"
)

type API struct {
	StormServer *stormserver.Client
}

func NewAPI(config ClientConfig) *API {
	// Initialize http backend
	client := client.NewClient(config)

	return &API{
		StormServer: &stormserver.Client{Client: client}
	}
}
