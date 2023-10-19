package webserver

import (
	"fmt"

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
}

func NewFiberWebServer(appConf *conf.AppConf) *FiberWebServer {
	return &FiberWebServer{
		Server: fiber.New(),
		Config: appConf,
	}
}

func (ws *FiberWebServer) StartServer() error {
	ws.Server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from go")
	})

	return ws.Server.Listen(fmt.Sprintf(":%d", ws.Config.Port))
}

func (ws *FiberWebServer) ShutdownServer() error {
	return ws.Server.Shutdown()
}
