package main

import (
	"context"

	"github.com/Mth-Ryan/waveaction/cmd/app/controllers"
	"github.com/Mth-Ryan/waveaction/cmd/app/webserver"
	"github.com/Mth-Ryan/waveaction/internal/conf"
	"go.uber.org/fx"
)

func RegisterWebServer(lc fx.Lifecycle, ws *webserver.FiberWebServer) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go ws.StartServer()
			return nil
		},
		OnStop: func(_ context.Context) error {
			go ws.ShutdownServer()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		conf.Module,
		controllers.Module,
		webserver.Module,
		fx.Invoke(RegisterWebServer),
	)
	app.Run()
}
