package networkzone

import (
	liquidweb "git.liquidweb.com/masre/liquidweb-go"
	"git.liquidweb.com/masre/liquidweb-go/client"
)

type NetworkZoneBackend interface {
	Details() (*liquidweb.Zone, *client.LWAPIError)
	List() (*[]liquidweb.Zone, *client.LWAPIError)
}

type Client struct {
	Backend *client.Client
}

func (c *Client) Details(id int) *liquidweb.ZoneItem {
	var zoneResult *liquidweb.ZoneItem
	zoneParams := liquidweb.ZoneParams{ID: id}

	c.Backend.CallInto("v1/Network/Zone/detail", zoneParams, zoneResult)

	return zoneResult
}

// List returns a list of network zones.
func (c *Client) List(params *liquidweb.ZoneListParams) (*liquidweb.ZoneList, error) {
	zoneList := &liquidweb.ZoneList{}

	err := c.Backend.CallInto("v1/Network/Zone/list", params, zoneList)
	if err != nil {
		return nil, err
	}

	return zoneList, nil
}
