package conf

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewAppConf))

type AppConf struct {
	RunMode string
	Port    int `config:"port"`
}

func NewAppConf() *AppConf {
	return getAppConf()
}

func getAppConf() *AppConf {
	appConf := &AppConf{}

	config.WithOptions(config.ParseEnv, func(o *config.Options) {
		o.DecoderConfig.TagName = "config"
	})

	config.AddDriver(yaml.Driver)

	confFilename := "./app-conf-dev.yml"
	if RunMode == RUN_MODE_RELEASE {
		confFilename = "./app-conf.yml"
	}

	if err := config.LoadFiles(confFilename); err != nil {
		panic(err)
	}

	if err := config.Decode(&appConf); err != nil {
		panic(err)
	}

	appConf.RunMode = RunMode
	return appConf
}
