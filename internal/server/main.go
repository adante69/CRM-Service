package server

import (
	"CRM-Service/internal/handlers"
	"CRM-Service/internal/middleware"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func NewHTTPServer(lc fx.Lifecycle, contactHandler *handlers.ContactHandler,
	partnerHandler *handlers.PartnerHandler, bidHandler *handlers.BidHandler,
	authHandler *handlers.AuthHandler) *http.Server {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/auth/signup", authHandler.SignUp).Methods("POST")
	r.HandleFunc("/auth/signin", authHandler.SignIn).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	// Contact routes
	api.HandleFunc("/contact", contactHandler.GetAllContacts).Methods("GET")
	api.HandleFunc("/contact/{id}", contactHandler.GetContact).Methods("GET")
	api.HandleFunc("/contact", contactHandler.CreateContact).Methods("POST")
	api.HandleFunc("/contact/{id}", contactHandler.UpdateContact).Methods("PUT")
	api.HandleFunc("/contact/{id}", contactHandler.DeleteContact).Methods("DELETE")
	// Partner routes
	api.HandleFunc("/partner", partnerHandler.GetAllPartners).Methods("GET")
	api.HandleFunc("/partner/{id}", partnerHandler.GetPartner).Methods("GET")
	api.HandleFunc("/partner", partnerHandler.CreatePartner).Methods("POST")
	api.HandleFunc("/partner/{id}", partnerHandler.UpdatePartner).Methods("PUT")
	api.HandleFunc("/partner/{id}", partnerHandler.DeletePartner).Methods("DELETE")
	// Bid routes
	api.HandleFunc("/bid", bidHandler.GetAllBids).Methods("GET")
	api.HandleFunc("/partner/{id}", bidHandler.GetBid).Methods("GET")
	api.HandleFunc("/bid", bidHandler.CreateBid).Methods("POST")
	api.HandleFunc("/bid/{id}", bidHandler.UpdateBid).Methods("PUT")
	api.HandleFunc("/bid/{id}", bidHandler.DeleteBid).Methods("DELETE")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
