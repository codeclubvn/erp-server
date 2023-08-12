package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pet-project/model"
)

type Migration struct {
	db *gorm.DB
}

type IMigration interface {
	Migrate(c *gin.Context)
}

func NewMigration(db *gorm.DB) *Migration {
	return &Migration{db: db}
}

func (h *Migration) Migrate(c *gin.Context) {
	if err := h.db.AutoMigrate(&model.User{}); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
}
