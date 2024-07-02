package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func NewMuxRouter() *mux.Router {
	return mux.NewRouter()
}

func NewHTTPServer(lc fx.Lifecycle, mux *mux.Router) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
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
