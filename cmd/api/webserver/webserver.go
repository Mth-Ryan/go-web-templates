package webserver

import (
	"fmt"

	"github.com/Mth-Ryan/waveaction/cmd/api/controllers"
	"github.com/Mth-Ryan/waveaction/internal/conf"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewFiberWebServer))

type WebServer interface {
	StartServer() error
	ShutdownServer() error
}

type FiberWebServer struct {
	Server *fiber.App
	Config *conf.AppConf
	Controllers []controllers.BaseController
}

func NewFiberWebServer(
	appConf *conf.AppConf,
	indexController *controllers.IndexController,
	booksController *controllers.BooksController,
) *FiberWebServer {
	return &FiberWebServer{
		Server: fiber.New(),
		Config: appConf,
		Controllers: []controllers.BaseController {
			indexController,
			booksController,
		},
	}
}

func (ws *FiberWebServer) StartServer() error {
	ws.Server.Use(logger.New())

	for _, controller := range ws.Controllers {
		_ = controller.GetRouter(ws.Server)
	}

	return ws.Server.Listen(fmt.Sprintf(":%d", ws.Config.Port))
}

func (ws *FiberWebServer) ShutdownServer() error {
	return ws.Server.Shutdown()
}
