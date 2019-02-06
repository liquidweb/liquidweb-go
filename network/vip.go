package network

import (
	liquidweb "git.liquidweb.com/masre/liquidweb-go"
)

// VIPParams is the set of parameters used when creating or updating a VIP.
type VIPParams struct {
	Domain     string `json:"domain,omitempty"`
	Zone       int    `json:"zone,omitempty"`
	UniqID     string `json:"uniq_id,omitempty"`
	SubAccount string `json:"subaccnt,omitempty"`
}

// VIP is the resource representing a VIP entry.
type VIP struct {
	Domain       string `json:"domain,omitempty"`
	Zone         int    `json:"zone,omitempty"`
	Active       bool   `json:"active,omitempty"`
	ActiveStatus string `json:"activeStatus,omitempty"`
	UniqID       string `json:"uniq_id,omitempty"`
	Destroyed    string `json:"destroyed,omitempty"`
}

// VIPItem is an envelope for the API result containing either a VIP or an error.
type VIPItem struct {
	liquidweb.LWAPIError
	VIP
}

// VIPDeletion represents the API result when deleting a VIP.
type VIPDeletion struct {
	liquidweb.LWAPIError
	Destroyed string `json:"destroyed"`
}
