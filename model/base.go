package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID       `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatorID uuid.UUID       `json:"creator_id" gorm:"column:creator_id"`
	UpdaterID uuid.UUID       `json:"updater_id"`
	CreatedAt time.Time       `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" sql:"index"  swaggertype:"string"`
}
