package services

import (
	"time"

	"github.com/Mth-Ryan/waveaction/pkg/application/dtos"
	"github.com/google/uuid"
)

type BooksService interface {
	GetAll() ([]dtos.BookOutputDto, error)
	Get(id uuid.UUID) (dtos.BookOutputDto, error)
	Create(input dtos.BookInputDto) (dtos.BookOutputDto, error)
	Update(id uuid.UUID, input dtos.BookInputDto) (dtos.BookOutputDto, error)
	Delete(id uuid.UUID) error
}

type ActualBooksService struct{}

func NewActualBooksService() *ActualBooksService {
	return &ActualBooksService{}
}

func (b *ActualBooksService) GetAll() ([]dtos.BookOutputDto, error) {
	return []dtos.BookOutputDto{
		{
			ID:        uuid.New(),
			Title:     "Game of Thrones",
			Author:    "J.R.R Martin",
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Title:     "Fire and Blood",
			Author:    "J.R.R Martin",
			CreatedAt: time.Now(),
		},
	}, nil
}

func (b *ActualBooksService) Get(id uuid.UUID) (dtos.BookOutputDto, error) {
	return dtos.BookOutputDto{
		ID:        uuid.New(),
		Title:     "Game of Thrones",
		Author:    "J.R.R Martin",
		CreatedAt: time.Now(),
	}, nil
}

func (b *ActualBooksService) Create(input dtos.BookInputDto) (dtos.BookOutputDto, error) {
	return dtos.BookOutputDto{
		ID:        uuid.New(),
		Title:     input.Title,
		Author:    input.Author,
		CreatedAt: time.Now(),
	}, nil
}

func (b *ActualBooksService) Update(id uuid.UUID, input dtos.BookInputDto) (dtos.BookOutputDto, error) {
	return dtos.BookOutputDto{
		ID:        id,
		Title:     input.Title,
		Author:    input.Author,
		CreatedAt: time.Now(),
	}, nil
}

func (b *ActualBooksService) Delete(id uuid.UUID) error {
	return nil
}
