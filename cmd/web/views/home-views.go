package views

import "github.com/flosch/pongo2/v6"

var (
	homeIndexTempl = getTmpl("./templates/home/index.tmpl.html")
)

type HomeViews struct {}

func NewHomeViews() *HomeViews {
	return &HomeViews{}
}

func (hv *HomeViews) Index() ([]byte, error) {
	return homeIndexTempl.ExecuteBytes(pongo2.Context{
		"title": "Home",
	})
}
