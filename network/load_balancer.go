package network

import (
	"net"

	liquidweb "git.liquidweb.com/masre/liquidweb-go"
	"git.liquidweb.com/masre/liquidweb-go/types"
)

// LoadBalancerNodeParams is the resource representing a load balancer node.
type LoadBalancerNodeParams struct {
	Domain string `json:"domain,omitempty"`
	IP     net.IP `json:"ip,omitempty"`
	UniqID string `json:"uniq_id,omitempty"`
}

// LoadBalancerNode is the resource representing a load balancer node.
type LoadBalancerNode struct {
	Domain string `json:"domain,omitempty"`
	IP     net.IP `json:"ip,omitempty"`
	UniqID string `json:"uniq_id,omitempty"`
}

// LoadBalancerServiceParams is the set of parameters used when adding services to a load balancer.
type LoadBalancerServiceParams struct {
	UniqID   string `json:"uniq_id,omitempty"`
	SrcPort  int    `json:"src_port,omitempty"`
	DestPort int    `json:"dest_port,omitempty"`
}

// LoadBalancerService is the resource representing a load balancer service.
type LoadBalancerService struct {
	UniqID   string        `json:"uniq_id,omitempty"`
	SrcPort  types.FlexInt `json:"src_port,omitempty"`
	DestPort types.FlexInt `json:"dest_port,omitempty"`
	Protocol string        `json:"protocol,omitempty"`
}

// LoadBalancerParams is the set of parameters used when creating or updating a load balancer.
type LoadBalancerParams struct {
	UniqID             string                      `json:"uniq_id,omitempty"`
	Name               string                      `json:"name,omitempty"`
	Nodes              []string                    `json:"nodes,omitempty"`
	Region             int                         `json:"region,omitempty"`
	Services           []LoadBalancerServiceParams `json:"services,omitempty"`
	SessionPersistence bool                        `json:"session_persistence,omitempty"`
	SSLCert            string                      `json:"ssl_cert,omitempty"`
	SSLIncludes        bool                        `json:"ssl_includes,omitempty"`
	SSLInt             string                      `json:"ssl_int,omitempty"`
	SSLKey             string                      `json:"ssl_key,omitempty"`
	SSLTermination     bool                        `json:"ssl_termination,omitempty"`
	Strategy           string                      `json:"strategy,omitempty"`
}

// LoadBalancer is the resource representing a load balancer.
type LoadBalancer struct {
	Name               string                 `json:"name,omitempty"`
	Nodes              []LoadBalancerNode     `json:"nodes,omitempty"`
	RegionID           types.FlexInt          `json:"region_id,omitempty"`
	Services           []LoadBalancerService  `json:"services,omitempty"`
	SessionPersistence types.NumericalBoolean `json:"session_persistence,omitempty"`
	SSLIncludes        types.NumericalBoolean `json:"ssl_includes,omitempty"`
	SSLTermination     types.NumericalBoolean `json:"ssl_termination,omitempty"`
	Strategy           string                 `json:"ip,omitempty"`
	UniqID             string                 `json:"uniq_id,omitempty"`
	VIP                string                 `json:"vip,omitempty"`
}

// LoadBalancerItem is an envelope for the API result containing either a load balancer or an error.
type LoadBalancerItem struct {
	liquidweb.LWAPIError
	LoadBalancer
}

// LoadBalancerDeletion represents the API result when deleting a load balancer.
type LoadBalancerDeletion struct {
	liquidweb.LWAPIError
	Destroyed string `json:"destroyed"`
}
