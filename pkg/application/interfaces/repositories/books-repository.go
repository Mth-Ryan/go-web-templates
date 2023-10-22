package repositories

import (
	"github.com/Mth-Ryan/waveaction/pkg/domain/entities"
	"github.com/google/uuid"
)

type BooksRepository interface {
	Get(id uuid.UUID) (entities.Book, error)
	GetAll() ([]entities.Book, error)
	Create(entity entities.Book) (entities.Book, error)
	Update(id uuid.UUID, entity entities.Book) (entities.Book, error)
	Delete(id uuid.UUID) error
}
