package storm

import (
	"net"
	"time"
)

// Server represents the underlying Storm VPS.
type Server struct {
	ACCNT               string
	Active              bool
	BackupEnabled       bool
	BackupPlan          string
	BackupQuota         int
	BackupSize          float64
	BandwidthQuota      int
	ConfigDescription   string
	ConfigID            int
	CreateDate          time.Time
	DiskSpace           int
	Domain              string
	IP                  net.IP
	IPCount             int
	ManageLevel         string
	Memory              int
	Template            string
	TemplateDescription string
	Type                string
	UniqID              string
	VCPU                int
	Zone                int
}

// ServerStates represents the various states the server can be in.
var ServerStates = []string{
	"Building",
	"Cloning",
	"Resizing",
	"Moving",
	"Booting",
	"Stopping",
	"Restarting",
	"Rebooting",
	"Shutting Down",
	"Restoring Backup",
	"Creating Image",
	"Deleting Image",
	"Restoring Image",
	"Re-Imaging",
	"Updating Firewall",
	"Updating Network",
	"Adding IPs",
	"Removing IP",
	"Destroying",
}
