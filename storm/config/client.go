package stormserver

import "git.liquidweb.com/masre/liquidweb-go/client"

type StormConfigClient interface {
	Details() (*StormConfig, *client.LWAPIError)
	List() (*[]StormConfig, *client.LWAPIError)
}

type Client struct {
	Client *client.Client
}

func (c *Client) Details() (*StormConfig, *client.LWAPIError) {
	return &StormConfig{}, nil
}

func (c *Client) List() (*StormConfig, *client.LWAPIError) {
	return &[]StormConfig{}, nil
}
