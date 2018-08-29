package client

import "git.liquidweb.com/masre/liquidweb-go/types"

type ListMeta struct {
	ItemCount types.FlexInt `json:"item_count,omitempty"`
	ItemTotal types.FlexInt `json:"item_total,omitempty"`
	PageNum   types.FlexInt `json:"page_num,omitempty"`
	PageSize  types.FlexInt `json:"page_size,omitempty"`
	PageTotal types.FlexInt `json:"page_total,omitempty"`
}
