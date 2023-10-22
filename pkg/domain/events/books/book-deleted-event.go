package books

import (
	"time"

	"github.com/Mth-Ryan/waveaction/pkg/domain/entities"
)

type BookDeletedEvent struct {
	EventKind int           `json:"event_kind"`
	OldBook   entities.Book `json:"old_book"`
	Timestamp time.Time     `json:"timestamp"`
}

func NewBookDeletedEvent(oldBook entities.Book) BookDeletedEvent {
	return BookDeletedEvent{
		EventKind: EVENT_KIND_DELETED,
		OldBook:   oldBook,
		Timestamp: time.Now(),
	}
}
