package webserver

import (
	"fmt"

	"github.com/Mth-Ryan/waveaction/cmd/app/controllers"
	"github.com/Mth-Ryan/waveaction/internal/conf"
	"github.com/gofiber/fiber/v2"
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
	homeController *controllers.HomeController,
	booksController *controllers.BooksController,
) *FiberWebServer {
	return &FiberWebServer{
		Server: fiber.New(),
		Config: appConf,
		Controllers: []controllers.BaseController {
			indexController,
			homeController,
			booksController,
		},
	}
}

func (ws *FiberWebServer) StartServer() error {
	for _, controller := range ws.Controllers {
		_ = controller.GetRouter(ws.Server)
	}

	return ws.Server.Listen(fmt.Sprintf(":%d", ws.Config.Port))
}

func (ws *FiberWebServer) ShutdownServer() error {
	return ws.Server.Shutdown()
}
