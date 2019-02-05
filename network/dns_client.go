package network

import (
	liquidweb "git.liquidweb.com/masre/liquidweb-go"
)

// DNSBackend describes the interface for interactions with the API.
type DNSBackend interface {
	Create(*DNSRecordParams) (*DNSRecordItem, error)
	Details(int) *DNSRecordItem
	List(*DNSRecordParams) *DNSRecordList
	Update(*DNSRecordParams) *DNSRecordItem
	Delete(*DNSRecordParams) *DNSRecordDeletion
}

// DNSClient is the backend implementation for interacting with DNS Records.
type DNSClient struct {
	Backend liquidweb.Backend
}

// Create creates a new DNS Record.
func (c *DNSClient) Create(params *DNSRecordParams) (*DNSRecordItem, error) {
	var result DNSRecordItem
	err := c.Backend.CallInto("v1/Network/DNS/Record/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details returns details about a DNS Record.
func (c *DNSClient) Details(id int) *DNSRecordItem {
	var result DNSRecordItem
	params := DNSRecordParams{ID: id}

	c.Backend.CallInto("v1/Network/DNS/Record/details", params, &result)

	return &result
}

// List returns a list of DNS Records.
func (c *DNSClient) List(params *DNSRecordParams) *DNSRecordList {
	list := &DNSRecordList{}

	c.Backend.CallInto("v1/Network/DNS/Record/list", params, list)
	return list
}

// Update will update a DNS Record.
func (c *DNSClient) Update(params *DNSRecordParams) *DNSRecordItem {
	var result DNSRecordItem
	c.Backend.CallInto("v1/Network/DNS/Record/update", params, &result)

	return &result
}

// Delete will delete a DNS Record.
func (c *DNSClient) Delete(params *DNSRecordParams) *DNSRecordDeletion {
	var result DNSRecordDeletion
	c.Backend.CallInto("v1/Network/DNS/Record/delete", params, &result)

	return &result
}
