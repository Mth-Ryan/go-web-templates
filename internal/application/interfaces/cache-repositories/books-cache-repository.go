package cacherepositories

import (
	"github.com/Mth-Ryan/waveaction/internal/domain/entities"
	"github.com/google/uuid"
)

type BooksCacheRepository interface {
	Set(entity entities.Book) error
	Get(id uuid.UUID) (entities.Book, error)
	Delete(id uuid.UUID) error
}
