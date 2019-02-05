package network

import (
	liquidweb "git.liquidweb.com/masre/liquidweb-go"
)

// VIPBackend describes the interface for interactions with the API.
type VIPBackend interface {
	Create(*VIPParams) (*VIPItem, error)
	Destroy(*VIPParams) *VIPDeletion
	Details(*VIPParams) *VIPItem
}

// VIPClient is the backend implementation for interacting with VIP.
type VIPClient struct {
	Backend liquidweb.Backend
}

// Create creates a new VIP.
func (c *VIPClient) Create(params *VIPParams) (*VIPItem, error) {
	var result VIPItem
	err := c.Backend.CallInto("v1/VIP/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details returns details about a VIP.
func (c *VIPClient) Details(uniqID string) *VIPItem {
	var result VIPItem
	params := VIPParams{UniqID: uniqID}

	c.Backend.CallInto("v1/VIP/details", params, &result)

	return &result
}

// Delete will delete a VIP.
func (c *VIPClient) Delete(params *VIPParams) *VIPDeletion {
	var result VIPDeletion
	c.Backend.CallInto("v1/VIP/destroy", params, &result)

	return &result
}
