package views

import (
	"io"

	"github.com/gofiber/fiber/v2"
)

type ViewsFiberAdapter struct {
	views ViewsRenderer
}

func (a *ViewsFiberAdapter) Load() error {
	return a.views.Load()
}

func (a *ViewsFiberAdapter) Render(
	writer io.Writer,
	template string,
	binding interface{},
	layouts...string,
) error {
	context := make(map[string]any)

	switch c := binding.(type) {
	case map[string]any:
		context = c
	case fiber.Map:
		context = c
	}

	return a.views.Render(writer, template, context)
}

func NewViewsFiberAdapter(views ViewsRenderer) fiber.Views {
	return &ViewsFiberAdapter{ views }
}
