package main

import (
	"context"
	nethttp "net/http"
	"os"
	"path"
	"time"

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

	commandController, err := http.NewCommandController(
		orderRepostiory, ep, time.Duration(cfg.Context.Timeout)*time.Second)
	must.NotFail(err)

	queryController := http.NewQueryController(service)

	e := echo.New()
	e.Logger.SetOutput(os.Stdout)

	srv := http.NewServer(cfg, queryController, commandController)

	go func() {
		if err := srv.Start(); err != nil && err != nethttp.ErrServerClosed {
			srv.Fatal(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Server.Timeout)*time.Second)
	defer cancel()

	graceful.Run()

	if err := srv.Shutdown(ctx); err != nil {
		srv.Fatal(err)
	}
}
