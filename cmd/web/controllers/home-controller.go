package controllers

import (
	"github.com/Mth-Ryan/waveaction/cmd/web/views"
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	views views.ViewsRenderer
}

func NewHomeController(views views.ViewsRenderer) *HomeController {
	return &HomeController {
		views,
	}
}

func (hc *HomeController) Index(ctx *fiber.Ctx) error {
	return renderView(
		ctx,
		hc.views,
		"./templates/home/index.tmpl.html",
		map[string]any{ 
			"title": "Home",
		},
	)
}

func (hc *HomeController) RegisterController(app *fiber.App) {
	router := app.Group("/home")
	
	router.Get("/", hc.Index)
}

