package main

import (
	"context"
	"github.com/core-go/health/server"
	"github.com/core-go/mq/config"
	"go-service/internal/app"
)

func main() {
	var conf app.Root
	er1 := config.Load(&conf, "configs/config")
	if er1 != nil {
		panic(er1)
	}
	ctx := context.Background()

	app, er2 := app.NewApp(ctx, conf)
	if er2 != nil {
		panic(er2)
	}

	go server.Serve(conf.Server, app.HealthHandler.Check)
	app.Receive(ctx, app.Handler.Handle)
}
