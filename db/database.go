package db

import (
	config2 "CRM-Service/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDataBase() (*gorm.DB, error) {
	config, err := config2.LoadConfiguration()
	if err != nil {
		panic(err)
	}
	dsn := config.Database.Dsn
	fmt.Println(dsn)
	fmt.Println(dsn)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}
