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

// Create will create a new storm server.
func (c *ServerClient) Create() *ServerItem {
	return &ServerItem{}
}

// Details fetches the details for a storm server.
func (c *ServerClient) Details() *ServerItem {
	return &ServerItem{}
}

// Update will update a storm server.
func (c *ServerClient) Update() *ServerItem {
	return &ServerItem{}
}

// Destroy will destroy a storm server.
func (c *ServerClient) Destroy(id string) *ServerItem {
	return &ServerItem{}
}
