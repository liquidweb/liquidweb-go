package storage

import (
	liquidweb "git.liquidweb.com/masre/liquidweb-go"
)

// BlockVolumeBackend describes the interface for interactions with the API.
type BlockVolumeBackend interface {
	Create(*BlockVolumeParams) (*BlockVolumeItem, error)
	Details(string) *BlockVolumeItem
	List(*BlockVolumeParams) *BlockVolumeList
	Update(*BlockVolumeParams) *BlockVolumeItem
	Delete(*BlockVolumeParams) *BlockVolumeDeletion
	Resize(*BlockVolumeParams) *BlockVolumeResize
}

// BlockVolumeClient is the backend implementation for interacting with block volumes.
type BlockVolumeClient struct {
	Backend liquidweb.Backend
}

// Create creates a new block volume.
func (c *BlockVolumeClient) Create(params *BlockVolumeParams) (*BlockVolumeItem, error) {
	var result BlockVolumeItem
	err := c.Backend.CallInto("v1/Storage/Block/Volume/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details returns details about a block volume.
func (c *BlockVolumeClient) Details(id string) *BlockVolumeItem {
	var result BlockVolumeItem
	params := BlockVolumeParams{UniqID: id}

	c.Backend.CallInto("v1/Storage/Block/Volume/details", params, &result)

	return &result
}

// List returns a list of block volumes.
func (c *BlockVolumeClient) List(params *BlockVolumeParams) *BlockVolumeList {
	list := &BlockVolumeList{}

	c.Backend.CallInto("v1/Storage/Block/Volume/list", params, list)
	return list
}

// Update will update a block volume.
func (c *BlockVolumeClient) Update(params *BlockVolumeParams) *BlockVolumeItem {
	var result BlockVolumeItem
	c.Backend.CallInto("v1/Storage/Block/Volume/update", params, &result)

	return &result
}

// Delete will delete a block volume.
func (c *BlockVolumeClient) Delete(params *BlockVolumeParams) *BlockVolumeDeletion {
	var result BlockVolumeDeletion
	c.Backend.CallInto("v1/Storage/Block/Volume/delete", params, &result)

	return &result
}

// Resize will resize a block volume.
func (c *BlockVolumeClient) Resize(params *BlockVolumeParams) *BlockVolumeResize {
	var result BlockVolumeResize
	c.Backend.CallInto("v1/Storage/Block/Volume/resize", params, &result)

	return &result
}
