package main

import (
	"CRM-Service/config"
	"CRM-Service/server"
	"fmt"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	configer, err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}
	fmt.Println("Environment:", configer.Env)
	fmt.Println("Server Host:", configer.Server.Host)
	fmt.Println("Server Port:", configer.Server.Port)
	fmt.Println("Database DSN:", configer.Database.Dsn)
	fmt.Println()
	fx.New(
		fx.Provide(server.NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()

}
