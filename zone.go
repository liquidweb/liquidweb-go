package liquidweb

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	"git.liquidweb.com/masre/liquidweb-go/types"
)

// Zone is a grouping of network resources.
type Zone struct {
	ID             int                               `json:"id,omitempty"`
	IsDefault      types.NumericalBoolean            `json:"is_default,omitempty"`
	Name           string                            `json:"name,omitempty"`
	Region         map[string]types.NumericalBoolean `json:"region,omitempty"`
	Status         string                            `json:"status,omitempty"`
	ValidSourceHVS map[string]types.NumericalBoolean `json:"valid_source_hvs,omitempty"`
}

// ZoneParams are the set of parameters you can pass to the API for Network Zones.
type ZoneParams struct {
	ID int
}

// ZoneList is an envelope for the API result containing either a list of zones or an error.
type ZoneList struct {
	*client.LWAPIError
	*client.ListMeta
	Items []Zone
}

// ZoneItem is an envelope for the API result containing either a zones or an error.
type ZoneItem struct {
	*client.LWAPIError
	Zone
}
