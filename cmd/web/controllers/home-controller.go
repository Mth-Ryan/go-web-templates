package controllers

import (
	"github.com/Mth-Ryan/waveaction/cmd/web/views"
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	views *views.HomeViews
}

func NewHomeController(views *views.HomeViews) *HomeController {
	return &HomeController {
		views,
	}
}

func (hc *HomeController) Index(ctx *fiber.Ctx) error {
	raw, err := hc.views.Index()
	if err != nil {
		return err
	}

	return renderView(ctx, raw)
}

func (hc *HomeController) RegisterController(app *fiber.App) {
	router := app.Group("/home")
	
	router.Get("/", hc.Index)
}

