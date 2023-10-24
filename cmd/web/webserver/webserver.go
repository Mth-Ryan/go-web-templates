package webserver

import (
	"fmt"

	"github.com/Mth-Ryan/go-web-templates/cmd/web/controllers"
	"github.com/Mth-Ryan/go-web-templates/pkg/conf"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewFiberWebServer,
		fx.ParamTags(`group:"controllers"`),
		fx.As(new(WebServer)),
	),
)

type WebServer interface {
	StartServer() error
	ShutdownServer() error
}

type FiberWebServer struct {
	Server *fiber.App
	Config *conf.AppConf
}

func NewFiberWebServer(
	controllers []controllers.BaseController, // This must be the first parameter
	appConf *conf.AppConf,
) *FiberWebServer {
	server := fiber.New()
	
	server.Use(logger.New())
	server.Static("/", "./public")

	for _, controller := range controllers {
		controller.RegisterController(server)
	}

	return &FiberWebServer{
		Server: server,
		Config: appConf,
	}
}

func (ws *FiberWebServer) StartServer() error {
	return ws.Server.Listen(fmt.Sprintf(":%d", ws.Config.Port))
}

func (ws *FiberWebServer) ShutdownServer() error {
	return ws.Server.Shutdown()
}
