package network

import (
	liquidweb "git.liquidweb.com/masre/liquidweb-go"
)

// VIPParams is the set of parameters used when creating or updating a VIP.
type VIPParams struct {
	Domain       string `json:"domain,omitempty"`
	Zone         string `json:"zone,omitempty"`
	Active       bool   `json:"active,omitempty"`
	ActiveStatus string `json:"activestatus,omitempty"`
	UniqID       string `json:"uniqID,omitempty"`
	SubAccount   string `json:"subaccount,omitempty"`
	Destroyed    string `json:"destroyed,omitempty"`
}

// VIP is the resource representing a VIP entry.
type VIP struct {
	Domain       string `json:"domain,omitempty"`
	Zone         string `json:"zone,omitempty"`
	Active       bool   `json:"active,omitempty"`
	ActiveStatus string `json:"activestatus,omitempty"`
	UniqID       string `json:"uniqID,omitempty"`
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
