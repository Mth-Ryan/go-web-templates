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

func (hc *HomeController) GetRouter(app *fiber.App) fiber.Router {
	router := app.Group("/home")
	
	router.Get("/", hc.Index)

	return router
}

