package storm

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
