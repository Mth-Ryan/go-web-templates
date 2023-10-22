package repositories

import (
	"github.com/Mth-Ryan/waveaction/pkg/application/interfaces/repositories"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewBooksRepository,
		fx.As(new(repositories.BooksRepository)),
	),
)
