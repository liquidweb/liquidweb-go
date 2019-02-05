package network

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
)

// ZoneBackend is the interface for network zones.
type ZoneBackend interface {
	Details(int) *ZoneItem
	List(*ZoneListParams) (*ZoneList, error)
}

// ZoneClient is the API client for network zones.
type ZoneClient struct {
	Backend *client.Client
}

// Details fetches the details for a zone.
func (c *ZoneClient) Details(id int) *ZoneItem {
	var zoneResult *ZoneItem
	zoneParams := ZoneParams{ID: id}

	c.Backend.CallInto("v1/Network/Zone/detail", zoneParams, zoneResult)

	return zoneResult
}

// List returns a list of network zones.
func (c *ZoneClient) List(params *ZoneListParams) (*ZoneList, error) {
	zoneList := &ZoneList{}

	err := c.Backend.CallInto("v1/Network/Zone/list", params, zoneList)
	if err != nil {
		return nil, err
	}

	return zoneList, nil
}
