package stormserver

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	"git.liquidweb.com/masre/liquidweb-go/storm"
)

type StormConfigClient interface {
	Details() (*storm.ConfigItem, *client.LWAPIError)
	List() (*[]storm.ConfigList, *client.LWAPIError)
}

type Client struct {
	Client *client.Client
}

func (c *Client) Details() *storm.ConfigItem {
	return &storm.ConfigItem{}
}

func (c *Client) List() *storm.ConfigList {
	return &storm.ConfigList{}
}
