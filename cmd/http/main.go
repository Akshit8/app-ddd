package main

import (
	"context"
	nethttp "net/http"
	"os"
	"path"
	"time"

	"github.com/Akshit8/app-ddd/internal/app"
	"github.com/Akshit8/app-ddd/internal/app/query"
	"github.com/Akshit8/app-ddd/internal/http"
	"github.com/Akshit8/app-ddd/internal/infra"
	"github.com/Akshit8/app-ddd/internal/infra/inmem"
	"github.com/Akshit8/app-ddd/pkg/graceful"
	"github.com/Akshit8/app-ddd/pkg/must"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	var cfg http.Config

	// get current working directory
	wd, _ := os.Getwd()

	viper.SetConfigFile(path.Join(wd, "cmd", "http", "config.yaml"))

	must.NotFailf(viper.ReadInConfig)
	must.NotFail(viper.Unmarshal(&cfg))

	orderRepostiory := inmem.NewOrderRepository()
	service := query.NewOrderQueryService(orderRepostiory)
	ep := infra.NewConsoleBus()

	app.SetupMediator(orderRepostiory, ep)

	commandController := http.CommandController{}
	queryController := http.NewQueryController(service)

	e := echo.New()
	e.Logger.SetOutput(os.Stdout)

	srv := http.NewServer(cfg, queryController, commandController)

	go func() {
		if err := srv.Start(); err != nil && err != nethttp.ErrServerClosed {
			srv.Fatal(err)
		}
	}()

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Server.Timeout)*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			srv.Fatal(err)
		}
	}()

	graceful.Run()
}
