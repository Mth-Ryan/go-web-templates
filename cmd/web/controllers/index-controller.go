package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type IndexController struct {}

func NewIndexController() *IndexController {
	return &IndexController {}
}

func (ic *IndexController) Index(ctx *fiber.Ctx) error {
	return ctx.Redirect("/home", 302)
}

func (ic *IndexController) RegisterController(app *fiber.App) {
	router := app.Group("/")
	
	router.Get("/", ic.Index)
}

