package main

import (
	"context"

	"github.com/Mth-Ryan/waveaction/cmd/app/controllers"
	"github.com/Mth-Ryan/waveaction/cmd/app/webserver"
	"github.com/Mth-Ryan/waveaction/internal/conf"
	"github.com/Mth-Ryan/waveaction/pkg/application/mappers"
	appservices "github.com/Mth-Ryan/waveaction/pkg/application/services"
	"github.com/Mth-Ryan/waveaction/pkg/infra/data"
	"github.com/Mth-Ryan/waveaction/pkg/infra/repositories"
	infraservices "github.com/Mth-Ryan/waveaction/pkg/infra/services"
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
		data.Module,
		repositories.Module,
		mappers.Module,
		infraservices.Module,
		appservices.Module,
		controllers.Module,
		webserver.Module,
		fx.Invoke(RegisterWebServer),
	)
	app.Run()
}
