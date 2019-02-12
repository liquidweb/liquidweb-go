package network

import (
	liquidweb "git.liquidweb.com/masre/liquidweb-go"
)

// LoadBalancerBackend describes the interface for interactions with the API.
type LoadBalancerBackend interface {
	Create(LoadBalancerParams) (*LoadBalancerItem, error)
	Details(string) *LoadBalancerItem
	Update(LoadBalancerParams) *LoadBalancerItem
	Delete(string) *LoadBalancerDeletion
}

// LoadBalancerClient is the backend implementation for interacting with the API.
type LoadBalancerClient struct {
	Backend liquidweb.Backend
}

// Create creates a new load balancer.
func (c *LoadBalancerClient) Create(params LoadBalancerParams) (*LoadBalancerItem, error) {
	var result LoadBalancerItem
	err := c.Backend.CallInto("v1/Network/LoadBalancer/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details returns details about a load balancer.
func (c *LoadBalancerClient) Details(uniqID string) *LoadBalancerItem {
	var result LoadBalancerItem
	params := LoadBalancerParams{UniqID: uniqID}

	c.Backend.CallInto("v1/Network/LoadBalancer/details", params, &result)

	return &result
}

// Update will update a load balancer.
func (c *LoadBalancerClient) Update(params LoadBalancerParams) *LoadBalancerItem {
	var result LoadBalancerItem
	c.Backend.CallInto("v1/Network/LoadBalancer/update", params, &result)

	return &result
}

// Delete will delete a load balancer.
func (c *LoadBalancerClient) Delete(uniqID string) *LoadBalancerDeletion {
	var result LoadBalancerDeletion
	params := LoadBalancerParams{UniqID: uniqID}

	c.Backend.CallInto("v1/Network/LoadBalancer/delete", params, &result)

	return &result
}
