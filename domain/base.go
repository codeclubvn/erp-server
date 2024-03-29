package domain

import (
	"gorm.io/gorm"
	"time"

	uuid "github.com/satori/go.uuid"
)

type BaseModel struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UpdaterID uuid.UUID       `json:"updater_id"`
	CreatedAt time.Time       `gorm:"column:created_at;type:timestamp;default:now();not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at;type:timestamp;default:now();not null" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	updater_id := tx.Statement.Context.Value("x-user-id")
	if updater_id != nil {
		b.UpdaterID = uuid.FromStringOrNil(updater_id.(string))
	}

	return nil
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	updater_id := tx.Statement.Context.Value("x-user-id")
	if updater_id != nil {
		b.UpdaterID = uuid.FromStringOrNil(updater_id.(string))
	}

	return nil
}

func (b *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	updater_id := tx.Statement.Context.Value("x-user-id")
	if updater_id != nil {
		b.UpdaterID = uuid.FromStringOrNil(updater_id.(string))
	}

	return nil
}
