package storm

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	"git.liquidweb.com/masre/liquidweb-go/storm"
)

type ServerClient interface {
	List() (*[]storm.Server, *client.LWAPIError)
	Details() (*storm.Server, *client.LWAPIError)
	Update() (*storm.Server, *client.LWAPIError)
	Destroy() (*storm.Server, *client.LWAPIError)
}

type Client struct {
	Backend *client.Client
}

func (c *Client) Create() *storm.ServerItem {
	return &storm.ServerItem{}
}

func (c *Client) Details() *storm.ServerItem {
	return &storm.ServerItem{}
}

func (c *Client) Update() *storm.ServerItem {
	return &storm.ServerItem{}
}

func (c *Client) Destroy() *storm.ServerItem {
	return &storm.ServerItem{}
}
