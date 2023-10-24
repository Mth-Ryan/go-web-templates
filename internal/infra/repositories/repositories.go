package repositories

import (
	"github.com/Mth-Ryan/go-web-templates/internal/application/interfaces/repositories"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewBooksRepository,
		fx.As(new(repositories.BooksRepository)),
	),
)
