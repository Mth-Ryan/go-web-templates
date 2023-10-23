package eventhandlers

import (
	events "github.com/Mth-Ryan/waveaction/internal/domain/events/books"
)

type BooksEventHandler interface {
	Handle(event events.BookEvent) 
}
