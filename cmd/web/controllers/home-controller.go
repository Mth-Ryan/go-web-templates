package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {}

func NewHomeController() *HomeController {
	return &HomeController {}
}

func (hc *HomeController) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(
		struct { 
			Message string
		}{ 
			Message: "See the example CRUD in /books",
		},
	)
}

func (hc *HomeController) RegisterController(app *fiber.App) {
	router := app.Group("/")
	
	router.Get("/", hc.Index)
}

