package request

type PageOptions struct {
	Page   int64  `form:"page" json:"page"`
	Limit  int64  `form:"limit" json:"limit"`
	Sort   string `form:"sort" json:"sort"`
	Search string `form:"search" json:"search"`
}
