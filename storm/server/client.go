package stormserver

import "git.liquidweb.com/masre/liquidweb-go/client"

type StormServerClient interface {
	Create() (*StormServer, *client.LWAPIError)
	Details() (*StormServer, *client.LWAPIError)
	Update() (*StormServer, *client.LWAPIError)
	Destroy() (*StormServer, *client.LWAPIError)
}

type Client struct {
	Client *client.Client
}

func (c *Client) Create() (*StormServer, *client.LWAPIError) {
	return &StormServer{}, nil
}

func (c *Client) Details() (*StormServer, *client.LWAPIError) {
	return &StormServer{}, nil
}

func (c *Client) Update() (*StormServer, *client.LWAPIError) {
	return &StormServer{}, nil
}

func (c *Client) Destroy() (*StormServer, *client.LWAPIError) {
	return &StormServer{}, nil
}
