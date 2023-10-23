package services

import (
	"log"

	"github.com/Mth-Ryan/waveaction/pkg/application/dtos"
	cacherepositories "github.com/Mth-Ryan/waveaction/pkg/application/interfaces/cache-repositories"
	eventhandlers "github.com/Mth-Ryan/waveaction/pkg/application/interfaces/event-handlers"
	"github.com/Mth-Ryan/waveaction/pkg/application/interfaces/repositories"
	"github.com/Mth-Ryan/waveaction/pkg/application/mappers"
	events "github.com/Mth-Ryan/waveaction/pkg/domain/events/books"
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
	repository      repositories.BooksRepository
	mapper          *mappers.BooksMapper
	cacheRepository cacherepositories.BooksCacheRepository
	eventsHandler   eventhandlers.BooksEventHandler
}

func NewActualBooksService(
	repository repositories.BooksRepository,
	mapper *mappers.BooksMapper,
	cacheRepository cacherepositories.BooksCacheRepository,
	eventsHandler eventhandlers.BooksEventHandler,
) *ActualBooksService {
	return &ActualBooksService{
		repository,
		mapper,
		cacheRepository,
		eventsHandler,
	}
}

func (b *ActualBooksService) GetAll() ([]dtos.BookOutputDto, error) {
	entities, err := b.repository.GetAll()
	outputs := b.mapper.OutputsFromEntities(&entities)

	return outputs, err
}

func (b *ActualBooksService) Get(id uuid.UUID) (dtos.BookOutputDto, error) {
	entity, err := b.cacheRepository.Get(id)
	if err != nil {
		log.Printf("Cache miss for: %s. err: '%s'", id.String(), err)
		entity, err  = b.repository.Get(id)
	}
	output := b.mapper.OutputFromEntity(&entity)

	return output, err
}

func (b *ActualBooksService) Create(input dtos.BookInputDto) (dtos.BookOutputDto, error) {
	entity := b.mapper.EntityFromInput(&input)
	err := b.repository.Create(&entity)
	output := b.mapper.OutputFromEntity(&entity)

	event := events.NewBookCreatedEvent(entity)
	b.eventsHandler.Handle(event)

	return output, err
}

func (b *ActualBooksService) Update(id uuid.UUID, input dtos.BookInputDto) (dtos.BookOutputDto, error) {
	entity := b.mapper.EntityFromInput(&input)
	err := b.repository.Update(id, &entity)
	output := b.mapper.OutputFromEntity(&entity)

	event := events.NewBookUpdatedEvent(entity)
	b.eventsHandler.Handle(event)

	return output, err
}

func (b *ActualBooksService) Delete(id uuid.UUID) error {
	event := events.NewBookDeletedEvent(id)
	b.eventsHandler.Handle(event)

	return b.repository.Delete(id)
}
