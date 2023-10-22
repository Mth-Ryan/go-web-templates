package books

import (
	"time"

	"github.com/Mth-Ryan/waveaction/pkg/domain/entities"
)

type BookUpdatedEvent struct {
	EventKind int           `json:"event_kind"`
	OldBook   entities.Book `json:"old_book"`
	NewBook   entities.Book `json:"new_book"`
	Timestamp time.Time     `json:"timestamp"`
}

func NewBookUpdatedEvent(newBook entities.Book, oldBook entities.Book) BookUpdatedEvent {
	return BookUpdatedEvent{
		EventKind: EVENT_KIND_UPDATED,
		NewBook:   newBook,
		OldBook:   oldBook,
		Timestamp: time.Now(),
	}
}
