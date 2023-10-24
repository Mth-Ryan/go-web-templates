package eventhandlers

import (
	events "github.com/Mth-Ryan/go-web-templates/internal/domain/events/books"
)

type BooksEventHandler interface {
	Handle(event events.BookEvent) 
}
