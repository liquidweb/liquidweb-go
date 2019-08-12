package storm

import liquidweb "git.liquidweb.com/masre/liquidweb-go"

// ServerBackend is the interface for storm servers.
type ServerBackend interface {
	List() *ServerItem
	Details() *ServerItem
	Update() *ServerItem
	Destroy(string) *ServerItem
}

// ServerClient is the API client for storm servers.
type ServerClient struct {
	Backend liquidweb.Backend
}

// List will fetch a list of storm servers.
func (c *ServerClient) List() *ServerItem {
	return &ServerItem{}
}

// Create a new storm server.
func (c *ServerClient) Create(params ServerParams) (*ServerItem, error) {
	var result ServerItem
	err := c.Backend.CallInto("v1/Storm/Server/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details fetches the details for a storm server.
func (c *ServerClient) Details(params ServerParams) (*ServerItem, error) {
	var result ServerItem

	err := c.Backend.CallInto("v1/Storm/Server/details", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Update a storm server.
func (c *ServerClient) Update(params ServerParams) (*ServerItem, error) {
	var result ServerItem

	err := c.Backend.CallInto("v1/Storm/Server/update", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Destroy a storm server.
func (c *ServerClient) Destroy(id string) (*ServerDeletion, error) {
	var result ServerDeletion
	params := ServerParams{UniqID: id}

	err := c.Backend.CallInto("v1/Storm/Server/destroy", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Status returns the current status of a storm server.
func (c *ServerClient) Status(id string) (*ServerStatus, error) {
	var result ServerStatus
	params := ServerParams{UniqID: id}

	err := c.Backend.CallInto("v1/Storm/Server/status", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
