package stormserver

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	"git.liquidweb.com/masre/liquidweb-go/storm"
)

type ConfigClient interface {
	Details() (*storm.ConfigItem, *client.LWAPIError)
	List() (*[]storm.ConfigList, *client.LWAPIError)
}

type Client struct {
	Backend *client.Client
}

func (c *Client) Details() *storm.ConfigItem {
	return &storm.ConfigItem{}
}

func (c *Client) List(params storm.ConfigListParams) (*storm.ConfigList, error) {
	configList := &storm.ConfigList{}

	err := c.Backend.CallInto("v1/Storm/Config/list", params, configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}
