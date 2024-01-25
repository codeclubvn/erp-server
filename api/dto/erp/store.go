package erpdto

import "erp/api/request"

type CreateStoreRequest struct {
	Name         string `json:"name" validate:"required"`
	Avatar       string `json:"avatar" validate:"required"`
	Thumbnail    string `json:"thumbnail" validate:"required"`
	Bio          string `json:"bio" validate:"required"`
	Domain       string `json:"domain" validate:"required"`
	BusinessType string `json:"business_type" validate:"required"`
	OpendAt      string `json:"opend_at" validate:"required"`
	ClosedAt     string `json:"closed_at" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Location     string `json:"location" validate:"required"`
}

type UpdateStoreRequest struct {
	CreateStoreRequest
}

type ListStoreRequest struct {
	request.PageOptions
}
