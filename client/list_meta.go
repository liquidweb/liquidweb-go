package client

type ListMeta struct {
	ItemCount int `json:"item_count,omitempty"`
	ItemTotal int `json:"item_total,omitempty"`
	PageNum   int `json:"page_num,omitempty"`
	PageSize  int `json:"page_size,omitempty"`
	PageTotal int `json:"page_total,omitempty"`
}
