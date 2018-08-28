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
	Client *client.Client
}

func (c *Client) Create() (*storm.Server, *client.LWAPIError) {
	return &storm.Server{}, nil
}

func (c *Client) Details() (*storm.Server, *client.LWAPIError) {
	return &storm.Server{}, nil
}

func (c *Client) Update() (*storm.Server, *client.LWAPIError) {
	return &storm.Server{}, nil
}

func (c *Client) Destroy() (*storm.Server, *client.LWAPIError) {
	return &storm.Server{}, nil
}
