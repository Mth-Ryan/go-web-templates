package entities

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `sql:"id"`
	Title     string    `sql:"title"`
	Author    string    `sql:"author"`
	CreatedAt time.Time `sql:"created_at"`
}
