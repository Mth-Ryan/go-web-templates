package services

import (
	"github.com/Mth-Ryan/waveaction/pkg/application/interfaces"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewPlaygroundJsonValidator,
		fx.As(new(interfaces.JsonValidator)),
	),
)
