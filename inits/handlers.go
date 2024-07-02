package inits

import (
	"CRM-Service/internal/handlers"
	"CRM-Service/internal/repositories"
	"CRM-Service/internal/server"
	"CRM-Service/internal/services"
	"CRM-Service/migrations/db"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
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
		server.NewHTTPServer),
)
