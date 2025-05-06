package database

import (
	"SysAgent/config"
	"SysAgent/internal/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	connString := config.ConfigViper.DataBase.StringConnection
	typeDatabase := config.ConfigViper.DataBase.TypeDatabase

	var db *gorm.DB
	var err error

	switch typeDatabase {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(connString), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("Error to connect to database SQLite: %w", err)
		}
	default:
		return nil, fmt.Errorf("Type of dabase not suported: %s", typeDatabase)
	}

	if err := db.AutoMigrate(&models.CollectMetric{}, &models.Agent{}, &models.Script{}); err != nil {
		return nil, fmt.Errorf("erro ao executar migração: %w", err)
	}
	return db, nil
}
