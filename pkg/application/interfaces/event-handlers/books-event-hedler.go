package eventhandlers

import (
	events "github.com/Mth-Ryan/waveaction/pkg/domain/events/books"
)

type BooksEventHandler interface {
	Handle(event events.BookEvent) 
}
