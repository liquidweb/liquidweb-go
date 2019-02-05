package storage

import (
	liquidweb "git.liquidweb.com/masre/liquidweb-go"
)

// Attachment represents the attachment details for a block volume.
type Attachment struct {
	Device   string `json:"device,omitempty"`
	Resource string `json:"resource,omitempty"`
}

// BlockVolumeParams is the set of parameters used when creating or updating a block volume
type BlockVolumeParams struct {
	Attach      string `json:"attach,omitempty"`
	CrossAttach bool   `json:"cross_attach,omitempty"`
	DetachFrom  string `json:"detach_from,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Region      int    `json:"region,omitempty"`
	Size        int    `json:"size,omitempty"`
	To          string `json:"to,omitempty"`
	UniqID      string `json:"uniq_id,omitempty"`
	Zone        int    `json:"zone,omitempty"`
	liquidweb.PageParams
}

// BlockVolume is the resource representing a block volume.
type BlockVolume struct {
	AttachedTo       Attachment `json:"attachedTo,omitempty"`
	CrossAttach      bool       `json:"cross_attach,omitempty"`
	Domain           string     `json:"domain,omitempty"`
	Label            string     `json:"label,omitempty"`
	Size             int        `json:"size,omitempty"`
	Status           string     `json:"status,omitempty"`
	UniqID           string     `json:"uniq_id,omitempty"`
	ZoneAvailability []int      `json:"zoneAvailability,omitempty"`
}

// BlockVolumeList is an envelope for the API result containing either a list of block volumes or an error.
type BlockVolumeList struct {
	liquidweb.LWAPIError
	liquidweb.ListMeta
	Items []BlockVolume `json:"items,omitempty"`
}

// BlockVolumeItem is an envelope for the API result containing either a block volume or an error.
type BlockVolumeItem struct {
	liquidweb.LWAPIError
	BlockVolume
}

// BlockVolumeDeletion represents the API result when deleting a block volume.
type BlockVolumeDeletion struct {
	liquidweb.LWAPIError
	Deleted int `json:"deleted"`
}
