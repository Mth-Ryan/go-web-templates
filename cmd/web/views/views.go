package views

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewPongo2ViewsRenderer,
		fx.As(new(ViewsRenderer)),
	),
)

type ViewsRenderer interface {
	Render(templateFile string, context map[string]any) ([]byte, error)
}

