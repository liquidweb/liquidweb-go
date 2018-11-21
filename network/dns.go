package network

import "git.liquidweb.com/masre/liquidweb-go/client"

// DNSRecordParams is the set of parameters used when creating or updating a DNS record.
type DNSRecordParams struct {
	ID              int
	Name            string
	Prio            int
	RData           string
	TTL             int
	Type            string
	Zone            string
	ZoneID          int
	AdminEmail      string
	Created         string
	Exchange        string
	Expiry          int
	FullData        string
	LastUpdated     string
	Minimum         int
	Nameserver      string
	Port            int
	RefreshInterval int
	RegionOverrides RegionOverrides
	Retry           int
	Serial          int
	Target          string
	Weight          int
}

// RegionOverrides contains region data.
type RegionOverrides struct {
	RData    string
	Region   string
	RegionID int
}

// DNSRecord is the resource representing a DNS record entry.
type DNSRecord struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	Prio            int             `json:"prio"`
	RData           string          `json:"rdata"`
	TTL             int             `json:"ttl"`
	Type            string          `json:"type"`
	Zone            string          `json:"zone"`
	ZoneID          int             `json:"zone_id"`
	AdminEmail      string          `json:"adminEmail"`
	Created         string          `json:"created"`
	Exchange        string          `json:"exchange"`
	Expiry          int             `json:"expiry"`
	FullData        string          `json:"fullData"`
	LastUpdated     string          `json:"last_updated"`
	Minimum         int             `json:"minimum"`
	Nameserver      string          `json:"nameserver"`
	Port            int             `json:"port"`
	RefreshInterval int             `json:"refreshInterval"`
	RegionOverrides RegionOverrides `json:"regionOverrides"`
	Retry           int             `json:"retry"`
	Serial          int             `json:"serial"`
	Target          string          `json:"target"`
	Weight          int             `json:"weight"`
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
