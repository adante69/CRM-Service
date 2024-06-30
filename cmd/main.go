package main

import (
	"CRM-Service/internal/handlers"
	"CRM-Service/internal/repositories"
	"CRM-Service/internal/server"
	"CRM-Service/internal/services"
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
		fx.Provide(repositories.NewAccountRepository),
		fx.Provide(services.NewAuthService),
		fx.Provide(handlers.NewAuthHandler),

		fx.Provide(repositories.NewContactRepository),
		fx.Provide(services.NewContactService),
		fx.Provide(handlers.NewContactHandler),

		fx.Provide(repositories.NewPartnerRepository),
		fx.Provide(services.NewPartnerService),
		fx.Provide(handlers.NewPartnerHandler),

		fx.Provide(repositories.NewBidRepository),
		fx.Provide(services.NewBidService),
		fx.Provide(handlers.NewBidHandler),

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
