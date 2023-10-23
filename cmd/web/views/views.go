package views

import (
	"github.com/flosch/pongo2/v6"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewHomeViews,
	NewBooksViews,
)

func getTmpl(filename string) *pongo2.Template {
	return pongo2.Must(pongo2.FromFile(filename))
}
