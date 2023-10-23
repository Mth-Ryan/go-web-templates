package views

import "github.com/flosch/pongo2/v6"

var (
	booksIndexTempl = getTmpl("./templates/books/index.tmpl.html")
)

type BooksViews struct {}

func NewBooksViews() *BooksViews {
	return &BooksViews{}
}

func (hv *BooksViews) Index() ([]byte, error) {
	return homeIndexTempl.ExecuteBytes(pongo2.Context{
		"title": "Books",
	})
}
