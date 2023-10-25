package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController {}
}

func (hc *HomeController) Index(ctx *fiber.Ctx) error {
	return ctx.Render(
		"home/index",
		map[string]any{ 
			"title": "Home",
		},
	)
}

func (hc *HomeController) RegisterController(app *fiber.App) {
	router := app.Group("/home")
	
	router.Get("/", hc.Index)
}

