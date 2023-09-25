package repository

import (
	"erp/infrastructure"

	"gorm.io/gorm"
)

type TX struct {
	db infrastructure.Database
}

type TxFn func(*TX) error

func WithTransaction(db *infrastructure.Database, fn TxFn) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return fn(&TX{
			db: infrastructure.Database{
				DB: tx,
			},
		})
	})
}

func GetTX(tx *TX, db infrastructure.Database) {
	if tx == nil {
		tx = &TX{db: db}
	}
}
