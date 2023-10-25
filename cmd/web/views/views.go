package views

import (
	"io"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewPongo2ViewsFactory,
		fx.As(new(ViewsFactory)),
	),
)

type ViewsFactory interface {
	GetRenderer(folder string, extension string) ViewsRenderer
}

type ViewsRenderer interface {
	Load() error
	Render(writer io.Writer, template string, context map[string]any) error
}

