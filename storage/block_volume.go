package network

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
)

// Attachment represents the attachment details for a block volume.
type Attachment struct {
	Device   string `json:"device,omitempty"`
	Resource string `json:"resource,omitempty"`
}

// BlockVolumeParams is the set of parameters used when creating or updating a block volume
type BlockVolumeParams struct {
	AttachedTo       Attachment `json:"attachedTo,omitempty"`
	CrossAttach      bool       `json:"cross_attach,omitempty"`
	Domain           string     `json:"domain,omitempty"`
	Label            string     `json:"label,omitempty"`
	Size             int        `json:"size,omitempty"`
	Status           string     `json:"status,omitempty"`
	UniqID           string     `json:"uniq_id,omitempty"`
	ZoneAvailability []int      `json:"zoneAvailability,omitempty"`
	DetachFrom       string     `json:"detach_from,omitempty"`
	client.PageParams
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

// DNSRecordList is an envelope for the API result containing either a list of DNS Records or an error.
type DNSRecordList struct {
	client.LWAPIError
	client.ListMeta
	Items []DNSRecord `json:"items,omitempty"`
}

// DNSRecordItem is an envelope for the API result containing either a DNS Record or an error.
type DNSRecordItem struct {
	client.LWAPIError
	DNSRecord
}

// DNSRecordDeletion represents the API result when deleting a DNS Record.
type DNSRecordDeletion struct {
	client.LWAPIError
	Deleted int `json:"deleted"`
}
