package infrastructure

import (
	"database/sql"
	config "erp/config"
	"erp/models"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	*gorm.DB
	logger *zap.Logger
}

func NewDatabase(config *config.Config, logger *zap.Logger) *Database {
	var err error
	var sqlDB *sql.DB

	logger.Info("Connecting to database...")
	gormDB, err := getDatabaseInstance(config)
	if err != nil {
		for i := 0; i < 5; i++ {
			gormDB, err = getDatabaseInstance(config)
			if err == nil {
				break
			}
		}
	}

	if err != nil {
		logger.Fatal("Database connection error", zap.Error(err))
	} else {
		logger.Info("Database connected")
	}

	db := &Database{gormDB, logger}

	db.RegisterTables()

	if err != nil {
		logger.Fatal("Database connection error", zap.Error(err))
	}
	sqlDB, err = db.DB.DB()
	if err != nil {
		logger.Fatal("sqlDB connection error", zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}

func getDatabaseInstance(config *config.Config) (db *gorm.DB, err error) {
	switch config.Database.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Database.Username,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect database: %w", err)
		}
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			config.Database.Host, config.Database.Username, config.Database.Password, config.Database.Name,
			config.Database.Port, config.Database.SSLMode, config.Database.TimeZone)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			PrepareStmt: true,
		})

		if err != nil {
			return nil, fmt.Errorf("failed to connect database: %w", err)
		}

		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	}
	return db, nil
}

func (d Database) RegisterTables() {
	err := d.DB.AutoMigrate(
		models.User{},
		models.Role{},
		models.Permission{},
		models.Store{},
		models.UserRole{},
		models.Category{},
		models.Product{},
		models.CategoryProduct{},
		models.Customer{},
		models.Revenue{},
		models.Order{},
		models.Debt{},
		models.OrderItem{},
		models.Promote{},
		models.PromoteUse{},
	)

	if err != nil {
		d.logger.Fatal("Database migration error", zap.Error(err))
		os.Exit(0)
	}
	d.logger.Info("Database migration success")
}
