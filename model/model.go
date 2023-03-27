package model

import (
	"github.com/kamalesh889/spendingCalculator/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DbConn *gorm.DB
}

func InitializeDB(conf *config.Config) (*Database, error) {

	db, err := gorm.Open(postgres.Open(conf.DbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}
