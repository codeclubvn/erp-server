package domain

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
)

type File struct {
	BaseModel
	FileName      string          `json:"file_name" gorm:"column:file_name;type:varchar(50);not null"`
	Path          string          `json:"path" gorm:"column:path;type:varchar(255);not null"`
	Size          int64           `json:"size" gorm:"column:size;type:bigint;not null"`
	ExtensionName string          `json:"type" gorm:"column:extension_name;type:varchar(10);not null"`
	Data          json.RawMessage `json:"domain" gorm:"column:data;type:jsonb;" swaggertype:"string"` // save domain flexibly
	UserId        uuid.UUID       `json:"user_id" gorm:"column:user_id;type:uuid"`
	User          User            `json:"user" gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (File) TableName() string {
	return "files"
}
