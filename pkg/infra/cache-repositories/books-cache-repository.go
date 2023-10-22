package cacherepositories

import (
	"context"

	"github.com/Mth-Ryan/waveaction/pkg/domain/entities"
	"github.com/Mth-Ryan/waveaction/pkg/infra/data"
	"github.com/google/uuid"
)

type BooksCacheRepository struct {
	cache *data.Cache
}

func NewBooksCacheRepository(cache *data.Cache) *BooksCacheRepository {
	return &BooksCacheRepository{
		cache,
	}
}

func (r *BooksCacheRepository) Set(entity entities.Book) error {
	ctx := context.Background()

	return r.cache.Ctx.Set(
		ctx,
		entity.ID.String(),
		mustSerializeToJson(entity),
		0,
	).Err()
}

func (r *BooksCacheRepository) Get(id uuid.UUID) (entities.Book, error) {
	ctx := context.Background()

	entity := entities.Book{}

	raw, err := r.cache.Ctx.Get(ctx, id.String()).Result()
	if err != nil {
		return entities.Book{}, err
	}

	mustDeserializeFromJson(raw, &entity)

	return entity, nil
}

func (r *BooksCacheRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()

	_, err := r.cache.Ctx.Del(ctx, id.String()).Result()

	return err
}
