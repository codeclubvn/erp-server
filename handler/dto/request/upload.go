package request

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
)

type UploadFileRequest struct {
	File   *multipart.FileHeader `json:"file" swaggerignore:"true" validation:"required"`
	UserId uuid.UUID             `json:"user_id" swaggerignore:"true"`
}

type UploadFileResponse struct {
	URL string `json:"url"`
}

type UpdateFileRequest struct {
	ID       string                `json:"id" form:"id" validate:"required"`
	File     *multipart.FileHeader `json:"file" swaggerignore:"true"`
	FileName string                `json:"file_name" form:"file_name"`
	Data     json.RawMessage       `json:"data,omitempty" swaggertype:"array,string"`
	UserId   uuid.UUID             `json:"user_id" swaggerignore:"true"`
}

type DeleteFileRequest struct {
	ID     string    `json:"id" form:"id" validate:"required"`
	UserId uuid.UUID `json:"user_id" swaggerignore:"true"`
}

type FileResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Status      bool      `json:"status"`
	NumberFile  int       `json:"number_product"`
}

type CreateFileRequest struct {
	FileName      string `json:"file_name" binding:"required"`
	Path          string `json:"path" binding:"required"`
	Size          int64  `json:"size" binding:"required"`
	ExtensionName string `json:"type" binding:"required"`
	Data          string `json:"domain" binding:"required"`
	UserId        string `json:"user_id" binding:"required"`
}

type UploadImageRequest struct {
	File *multipart.FileHeader
}
