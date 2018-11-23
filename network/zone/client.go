package network

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	network "git.liquidweb.com/masre/liquidweb-go/network"
)

type NetworkZoneBackend interface {
	Details() (*network.Zone, *client.LWAPIError)
	List() (*[]network.Zone, *client.LWAPIError)
}

type Client struct {
	Backend *client.Client
}

func (c *Client) Details(id int) *network.ZoneItem {
	var zoneResult *network.ZoneItem
	zoneParams := network.ZoneParams{ID: id}

	c.Backend.CallInto("v1/Network/Zone/detail", zoneParams, zoneResult)

	return zoneResult
}

// List returns a list of network zones.
func (c *Client) List(params *network.ZoneListParams) (*network.ZoneList, error) {
	zoneList := &network.ZoneList{}

	err := c.Backend.CallInto("v1/Network/Zone/list", params, zoneList)
	if err != nil {
		return nil, err
	}

	return zoneList, nil
}
