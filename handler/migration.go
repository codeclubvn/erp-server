package handler

import (
	"erp-server/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
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

func (h *Migration) BaseMigrate(ctx *gin.Context) {
	if err := h.db.Exec(`
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		`).Error; err != nil {
		fmt.Errorf(err.Error())
	}
}

func (h *Migration) Migrate(ctx *gin.Context) {
	// put your migrations at the end of the list
	migrate := gormigrate.New(h.db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// add uuid extension
			// add table user, business, product, order, money
			ID: "20230814200514",
			Migrate: func(tx *gorm.DB) error {
				h.BaseMigrate(ctx)
				if err := h.db.AutoMigrate(
					&model.User{},
					&model.Business{},
					&model.Product{},
					&model.Order{},
					&model.Money{}); err != nil {
					return err
				}
				return nil
			},
		},
	})
	err := migrate.Migrate()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}
