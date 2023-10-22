package cacherepositories

import (
	"encoding/json"

	cacherepositories "github.com/Mth-Ryan/waveaction/pkg/application/interfaces/cache-repositories"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewBooksCacheRepository,
		fx.As(new(cacherepositories.BooksCacheRepository)),
	),
)

func mustSerializeToJson(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if (err != nil) {
		panic(err)
	}

	return string(jsonData)
}


func mustDeserializeFromJson(data string, target interface{}) {
	err := json.Unmarshal([]byte(data), target)
	if err != nil {
		panic(err)
	}
}
