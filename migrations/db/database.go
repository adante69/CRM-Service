package db

import (
	"CRM-Service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDataBase() (*gorm.DB, error) {
	configer, err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}
	dsn := configer.Database.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}
