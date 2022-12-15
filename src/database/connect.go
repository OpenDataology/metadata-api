package database

import (
	"github.com/dataset-license/portal-backend/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	conf := config.Get()
	db, err := gorm.Open(postgres.Open(conf.DSN), &gorm.Config{})

	if err == nil {
		DB = db
		return db, err
	}
	return nil, err
}
