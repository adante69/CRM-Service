package db

import (
	"CRM-Service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDataBase(cf *config.Configuration) (*gorm.DB, error) {
	dsn := cf.Database.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}
