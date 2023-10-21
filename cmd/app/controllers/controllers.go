package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewControllersGroup)

func NewControllersGroup() *ControllersGroup {
	controllers := []BaseController {
		NewIndexController(),
		NewHomeController(),
	}

	return &ControllersGroup { controllers }
}

type BaseController interface {
	GetRouter(app *fiber.App) fiber.Router
}

type ControllersGroup struct {
	controllers []BaseController
}

func (c *ControllersGroup) GetAll() []BaseController {
	return c.controllers
}
