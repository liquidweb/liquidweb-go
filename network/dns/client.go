package network

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	network "git.liquidweb.com/masre/liquidweb-go/network"
)

// DNSRecordBackend describes the interface for interactions with the API.
type DNSRecordBackend interface {
	Create(*network.DNSRecordParams) *network.DNSRecordItem
	Details(int) *network.DNSRecordItem
	List(*network.DNSRecordParams) *network.DNSRecordList
}

// Client is the backend implementation for interacting with DNS Records.
type Client struct {
	Backend *client.Client
}

// Create creates a new DNS Record.
func (c *Client) Create(params *network.DNSRecordParams) *network.DNSRecordItem {
	var result *network.DNSRecordItem
	c.Backend.CallInto("v1/Network/DNS/Record/create", params, result)

	return result
}

// Details returns details about a DNS Record.
func (c *Client) Details(id int) *network.DNSRecordItem {
	var result *network.DNSRecordItem
	params := network.DNSRecordParams{ID: id}

	c.Backend.CallInto("v1/Network/DNS/Record/details", params, result)

	return result
}

// List returns a list of DNS Records.
func (c *Client) List(params *network.DNSRecordParams) *network.DNSRecordList {
	list := &network.DNSRecordList{}

	c.Backend.CallInto("v1/Network/DNS/Record/list", params, list)
	return list
}
