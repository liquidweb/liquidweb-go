package network

import "git.liquidweb.com/masre/liquidweb-go/types"

// Zone is a grouping of network resources.
type Zone struct {
	ID             int
	IsDefault      bool
	Name           string
	Region         *map[string]bool
	Status         string
	ValidSourceHVS *map[string]types.NumericalBoolean
}
