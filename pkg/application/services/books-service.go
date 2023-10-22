package services

import (
	"github.com/Mth-Ryan/waveaction/pkg/application/dtos"
	"github.com/Mth-Ryan/waveaction/pkg/application/interfaces/repositories"
	"github.com/Mth-Ryan/waveaction/pkg/application/mappers"
	"github.com/google/uuid"
)

type BooksService interface {
	GetAll() ([]dtos.BookOutputDto, error)
	Get(id uuid.UUID) (dtos.BookOutputDto, error)
	Create(input dtos.BookInputDto) (dtos.BookOutputDto, error)
	Update(id uuid.UUID, input dtos.BookInputDto) (dtos.BookOutputDto, error)
	Delete(id uuid.UUID) error
}

type ActualBooksService struct {
	repository repositories.BooksRepository
	mapper     *mappers.BooksMapper
}

func NewActualBooksService(
	repository repositories.BooksRepository,
	mapper *mappers.BooksMapper,
) *ActualBooksService {
	return &ActualBooksService{
		repository,
		mapper,
	}
}

func (b *ActualBooksService) GetAll() ([]dtos.BookOutputDto, error) {
	entities, err := b.repository.GetAll()
	outputs := b.mapper.OutputsFromEntities(&entities)

	return outputs, err
}

func (b *ActualBooksService) Get(id uuid.UUID) (dtos.BookOutputDto, error) {
	entity, err := b.repository.Get(id)
	output := b.mapper.OutputFromEntity(&entity)

	return output, err
}

func (b *ActualBooksService) Create(input dtos.BookInputDto) (dtos.BookOutputDto, error) {
	entity := b.mapper.EntityFromInput(&input)
	newEntity, err := b.repository.Create(entity)
	output := b.mapper.OutputFromEntity(&newEntity)

	return output, err
}

func (b *ActualBooksService) Update(id uuid.UUID, input dtos.BookInputDto) (dtos.BookOutputDto, error) {
	entity := b.mapper.EntityFromInput(&input)
	newEntity, err := b.repository.Update(id, entity)
	output := b.mapper.OutputFromEntity(&newEntity)

	return output, err
}

func (b *ActualBooksService) Delete(id uuid.UUID) error {
	return b.repository.Delete(id)
}
