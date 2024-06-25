package main

import (
	"CRM-Service/config"
	"CRM-Service/migrations/db"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
)

func main() {
	fx.New(
		fx.Provide(config.LoadConfiguration),
		fx.Provide(db.CreateDataBase),
		fx.Invoke(runMigrations),
	).Run()

}

func runMigrations(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}

	goose.SetDialect("postgres")
	if err := goose.Up(sqlDB, "./migrations"); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
}
