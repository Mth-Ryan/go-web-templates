package main

import (
	"context"

	"github.com/Mth-Ryan/go-web-templates/cmd/web/controllers"
	"github.com/Mth-Ryan/go-web-templates/cmd/web/views"
	"github.com/Mth-Ryan/go-web-templates/cmd/web/webserver"
	"github.com/Mth-Ryan/go-web-templates/internal/application/mappers"
	appservices "github.com/Mth-Ryan/go-web-templates/internal/application/services"
	cacherepositories "github.com/Mth-Ryan/go-web-templates/internal/infra/cache-repositories"
	"github.com/Mth-Ryan/go-web-templates/internal/infra/data"
	eventhandlers "github.com/Mth-Ryan/go-web-templates/internal/infra/event-handlers"
	"github.com/Mth-Ryan/go-web-templates/internal/infra/repositories"
	infraservices "github.com/Mth-Ryan/go-web-templates/internal/infra/services"
	"github.com/Mth-Ryan/go-web-templates/pkg/conf"
	"github.com/Mth-Ryan/go-web-templates/pkg/logger"
	"go.uber.org/fx"
)

func RegisterWebServer(lc fx.Lifecycle, ws webserver.WebServer) {
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
		views.Module,
		controllers.Module,
		webserver.Module,
		fx.Invoke(RegisterWebServer),
	)
	app.Run()
}
