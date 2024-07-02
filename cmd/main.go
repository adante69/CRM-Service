package main

import (
	"CRM-Service/inits"
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
		app := fx.New(
			inits.Modules,
			fx.Invoke(runMigrations),
		)
		app.Run()
		return
	}

	app := fx.New(
		inits.Modules,
		fx.Invoke(startApplication),
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
