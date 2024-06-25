package main

import (
	"CRM-Service/internal/server"
	"CRM-Service/migrations/db"
	"fmt"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(db.CreateDataBase),
		fx.Provide(server.NewHTTPServer),
		fx.Invoke(runMigrations),
		fx.Invoke(startApplication),
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

func startApplication(httpServer *http.Server) {
	fmt.Println("rabotaem")
}
