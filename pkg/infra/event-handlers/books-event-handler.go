package eventhandlers

import (
	"log"

	cacherepositories "github.com/Mth-Ryan/waveaction/pkg/application/interfaces/cache-repositories"
	events "github.com/Mth-Ryan/waveaction/pkg/domain/events/books"
)

type BooksEventHandler struct {
	cache cacherepositories.BooksCacheRepository
}

func NewBooksEventHandler(
	cache cacherepositories.BooksCacheRepository,
) *BooksEventHandler {
	return &BooksEventHandler{
		cache,
	}
}

// The ideal scenario is make this a async job with a queue
func (h *BooksEventHandler) Handle(event events.BookEvent) {
	switch e := event.(type) {

	case events.BookCreatedEvent:
		err := h.cache.Set(e.NewBook)
		if (err != nil) {
			log.Println(err)
		}

	case events.BookUpdatedEvent:
		err := h.cache.Set(e.NewBook)
		if (err != nil) {
			log.Println(err)
		}

	case events.BookDeletedEvent:
		err := h.cache.Delete(e.OldBookID)
		if (err != nil) {
			log.Println(err)
		}

	}
}
