package response

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SimpleResponse struct {
	//Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SimpleResponseList struct {
	//Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    Meta        `json:"meta"`
}

type Meta struct {
	Total       *int64 `json:"total"`
	Page        int64  `form:"page" json:"page"`
	Limit       int64  `form:"limit" json:"limit"`
	Sort        string `form:"sort" json:"sort"`
	PageCount   int64  `json:"page_count"`
	HasPrevPage bool   `json:"has_prev_page"`
	HasNextPage bool   `json:"has_next_page"`
}

type GetByIDsRequest struct {
	IDs []int64 `json:"ids"`
}
