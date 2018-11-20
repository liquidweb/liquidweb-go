package liquidweb

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	networkdns "git.liquidweb.com/masre/liquidweb-go/network/dns"
	networkzone "git.liquidweb.com/masre/liquidweb-go/network/zone"
	stormconfig "git.liquidweb.com/masre/liquidweb-go/storm/config"
	stormserver "git.liquidweb.com/masre/liquidweb-go/storm/server"
)

// API is the structure that houses all of our various API clients that interact with various Storm resources.
type API struct {
	NetworkDNS  networkdns.DNSRecordBackend
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
		NetworkDNS:  &networkdns.Client{Backend: client},
		NetworkZone: &networkzone.Client{Backend: client},
		StormConfig: &stormconfig.Client{Backend: client},
		StormServer: &stormserver.Client{Backend: client},
	}

	return api, nil
}
