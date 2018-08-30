package storm

import (
	"git.liquidweb.com/masre/liquidweb-go/client"
	"git.liquidweb.com/masre/liquidweb-go/types"
)

// Config represents the configuration of the server.
type Config struct {
	ID               types.FlexInt                     `json:"id"`
	Active           types.NumericalBoolean            `json:"active,omitempty"`
	Available        types.NumericalBoolean            `json:"available,omitempty"`
	Category         string                            `json:"category,omitempty"`
	Description      string                            `json:"description,omitempty"`
	Disk             types.FlexInt                     `json:"disk,omitempty"`
	Featured         types.NumericalBoolean            `json:"featured,omitempty"`
	Memory           types.FlexInt                     `json:"memory,omitempty"`
	VCPU             types.FlexInt                     `json:"vcpu,omitempty"`
	ZoneAvailability map[string]types.NumericalBoolean `json:"zone_availability,omitempty"`
}

// ConfigList is an envelope for the API result containing either a list of storm configs or an error.
type ConfigList struct {
	client.LWAPIError
	client.ListMeta
	Items []Config
}

// ConfigItem is an envelope for the API result containing either a storm config or an error.
type ConfigItem struct {
	client.LWAPIError
	Config
}

// ConfigListParams are the set of parameters you can pass to the API for listing storm configs.
type ConfigListParams struct {
	Available bool   `json:"available,omitempty"`
	Category  string `json:"category,omitempty"`
	PageNum   int    `json:"page_num,omitempty"`
	PageSize  int    `json:"page_size,omitempty"`
}
