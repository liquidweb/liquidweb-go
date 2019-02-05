package client

import "git.liquidweb.com/masre/liquidweb-go/types"

// PageParams support pagination parameters in parameter types.
type PageParams struct {
	PageNum  types.FlexInt `json:"page_num,omitempty"`
	PageSize types.FlexInt `json:"page_size,omitempty"`
}
