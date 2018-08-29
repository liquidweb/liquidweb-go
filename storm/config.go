package storm

import "git.liquidweb.com/masre/liquidweb-go/client"

// Config represents the configuration of the server.
type Config struct {
	ID                int
	Active            bool
	Available         bool
	category          string
	CPUCores          int
	CPUCount          int
	CPUHyperthreading bool
	CPUModel          string
	CPUSpeed          int
	Description       string
	Disk              int
	DiskCount         int
	DiskTotal         int
	DiskType          string
	Featured          bool
	Memory            int
	RaidLevel         int
	RAMAvailable      int
	RAMTotal          int
	VCPU              int
	ZoneAvailability  *map[string]bool
}

// ConfigList is an envelope for the API result containing either a list of storm configs or an error.
type ConfigList struct {
	*client.LWAPIError
	*client.ListMeta
	Items []Config
}

// ConfigItem is an envelope for the API result containing either a storm config or an error.
type ConfigItem struct {
	*client.LWAPIError
	Config
}
