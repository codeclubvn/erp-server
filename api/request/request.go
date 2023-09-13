package request

type PageOptions struct {
	Page   int    `form:"page" json:"page"`
	Limit  int    `form:"limit" json:"limit"`
	Sort   string `form:"sort" json:"sort"`
	Search string `form:"search" json:"search"`
}

func (p *PageOptions) GetOffset() int {
	return (p.Page - 1) * p.Limit
}
