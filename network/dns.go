package network

import "git.liquidweb.com/masre/liquidweb-go/client"

// DNSRecordParams is the set of parameters used when creating or updating a DNS record.
type DNSRecordParams struct {
	ID     int
	Name   string
	Prio   int
	RData  string
	TTL    int
	Type   string
	Zone   string
	ZoneID int
}

// DNSRecord is the resource representing a DNS record entry.
type DNSRecord struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Prio   int    `json:"prio"`
	RData  string `json:"rdata"`
	TTL    int    `json:"ttl"`
	Type   string `json:"type"`
	Zone   string `json:"zone"`
	ZoneID int    `json:"zone_id"`
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
