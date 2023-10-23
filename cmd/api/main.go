package main

import (
	"context"

	"github.com/Mth-Ryan/waveaction/cmd/api/controllers"
	"github.com/Mth-Ryan/waveaction/cmd/api/webserver"
	"github.com/Mth-Ryan/waveaction/pkg/conf"
	"github.com/Mth-Ryan/waveaction/pkg/logger"
	"github.com/Mth-Ryan/waveaction/internal/application/mappers"
	appservices "github.com/Mth-Ryan/waveaction/internal/application/services"
	cacherepositories "github.com/Mth-Ryan/waveaction/internal/infra/cache-repositories"
	"github.com/Mth-Ryan/waveaction/internal/infra/data"
	eventhandlers "github.com/Mth-Ryan/waveaction/internal/infra/event-handlers"
	"github.com/Mth-Ryan/waveaction/internal/infra/repositories"
	infraservices "github.com/Mth-Ryan/waveaction/internal/infra/services"
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
		logger.Module,
		data.Module,
		repositories.Module,
		cacherepositories.Module,
		mappers.Module,
		infraservices.Module,
		eventhandlers.Module,
		appservices.Module,
		controllers.Module,
		webserver.Module,
		fx.Invoke(RegisterWebServer),
	)
	app.Run()
}
