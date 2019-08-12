package storm

import (
	"net"
	"time"

	liquidweb "git.liquidweb.com/masre/liquidweb-go"
	"git.liquidweb.com/masre/liquidweb-go/types"
)

// ServerListParams is the set parameters used when listing Storm Servers.
type ServerListParams struct {
	liquidweb.PageParams
	Parent string `json:"parent,omitempty"`
}

// ServerParams is the set of parameters used when creating or updating a Storm Server.
type ServerParams struct {
	UniqID         string `json:"uniq_id,omitempty"`
	BackupEnabled  int    `json:"backup_enabled,omitempty"`
	BackupID       int    `json:"backup_id,omitempty"`
	BackupPlan     string `json:"backup_plan,omitempty"`
	BackupQuota    int    `json:"backup_quota,omitempty"`
	BandwidthQuota string `json:"bandwidth_quota,omitempty"`
	ConfigID       int    `json:"config_id,omitempty"`
	Domain         string `json:"domain,omitempty"`
	ImageID        int    `json:"image_id,omitempty"`
	IPCount        int    `json:"ip_count,omitempty"`
	MSSQL          string `json:"ms_sql,omitempty"`
	Password       string `json:"password,omitempty"`
	PublicSSHKey   string `json:"public_ssh_key,omitempty"`
	Template       string `json:"template,omitempty"`
	Zone           int    `json:"zone,omitempty"`
}

// Server represents the underlying Storm VPS.
type Server struct {
	ACCNT               string        `json:"accnt,omitempty"`
	Active              types.FlexInt `json:"active,omitempty"`
	BackupEnabled       types.FlexInt `json:"backup_enabled,omitempty"`
	BackupPlan          string        `json:"backup_plan,omitempty"`
	BackupQuota         string        `json:"backup_quota,omitempty"`
	BackupSize          float64       `json:"backup_size,omitempty"`
	BandwidthQuota      types.FlexInt `json:"bandwidth_quota,omitempty"`
	ConfigDescription   string        `json:"config_description,omitempty"`
	ConfigID            types.FlexInt `json:"config_id,omitempty"`
	CreateDate          time.Time     `json:"create_date,omitempty"`
	DiskSpace           types.FlexInt `json:"disk_space,omitempty"`
	Domain              string        `json:"domain,omitempty"`
	IP                  net.IP        `json:"ip,omitempty"`
	IPCount             types.FlexInt `json:"ip_count,omitempty"`
	ManageLevel         string        `json:"manage_level,omitempty"`
	Memory              types.FlexInt `json:"memory,omitempty"`
	Template            string        `json:"template,omitempty"`
	TemplateDescription string        `json:"template_description,omitempty"`
	Type                string        `json:"type,omitempty"`
	UniqID              string        `json:"uniq_id,omitempty"`
	VCPU                types.FlexInt `json:"vcpu,omitempty"`
	Zone                types.FlexInt `json:"zone,omitempty"`
}

// ServerStatus represents status of a Storm Server.
type ServerStatus struct {
	liquidweb.LWAPIError
	DetailedStatus string              `json:"detailed_status,omitempty"`
	Progress       types.FlexInt       `json:"progress,omitempty"`
	Running        ServerRunningStatus `json:"running,omitempty"`
	Status         string              `json:"status,omitempty"`
}

// ServerRunningStatus represents a detailed status step of a Storm Server.
type ServerRunningStatus struct {
	CurrentStep    string `json:"current_step,omitempty"`
	DetailedStatus string `json:"detailed_status,omitempty"`
	Name           string `json:"name,omitempty"`
	Status         string `json:"status,omitempty"`
}

// ServerDeletion represents the API result when deleting a Storm Server.
type ServerDeletion struct {
	liquidweb.LWAPIError
	Destroyed string `json:"destroyed"`
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

// ServerList is an envelope for the API result containing either a list of storm configs or an error.
type ServerList struct {
	liquidweb.LWAPIError
	liquidweb.ListMeta
	Items []Server
}

// ServerItem is an envelope for the API result containing either a storm server or an error.
type ServerItem struct {
	*liquidweb.LWAPIError
	Server
}
