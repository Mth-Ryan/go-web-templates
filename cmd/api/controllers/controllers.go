package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewIndexController,
	NewBooksController,
)

type BaseController interface {
	GetRouter(app *fiber.App) fiber.Router
}

