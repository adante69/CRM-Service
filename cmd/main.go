package main

import (
	"CRM-Service/config"
	"CRM-Service/internal/handlers"
	"CRM-Service/internal/repositories"
	"CRM-Service/internal/server"
	"CRM-Service/internal/services"
	"CRM-Service/migrations/db"
	"flag"
	"fmt"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	if *migrate {
		app := fx.New(fx.Invoke(runMigrations))
		app.Run()
		return
	}

	app := fx.New(
		fx.Provide(
			config.LoadConfiguration,
			db.CreateDataBase,
			repositories.NewAccountRepository,
			services.NewAuthService,
			handlers.NewAuthHandler,
			repositories.NewContactRepository,
			services.NewContactService,
			handlers.NewContactHandler,
			repositories.NewPartnerRepository,
			services.NewPartnerService,
			handlers.NewPartnerHandler,
			repositories.NewBidRepository,
			services.NewBidService,
			handlers.NewBidHandler,
			server.NewMuxRouter,
			server.NewHTTPServer),
		fx.Invoke(handlers.RegisterAuthRoutes,
			handlers.RegisterBidRoutes,
			handlers.RegisterContactsRoutes,
			handlers.RegisterPartnerRoutes,
			startApplication),
	)
	app.Run()
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
	fmt.Println("succesfully started application")
}
