package views

import (
	"github.com/Mth-Ryan/go-web-templates/pkg/conf"
	"github.com/Mth-Ryan/go-web-templates/pkg/logger"
	"github.com/flosch/pongo2/v6"
)

type Pongo2ViewsRenderer struct {
	appConf       *conf.AppConf
	logger        logger.ApplicationLogger
	templateCache map[string](*pongo2.Template)
}

func NewPongo2ViewsRenderer(
	appConf *conf.AppConf,
	logger logger.ApplicationLogger,
) *Pongo2ViewsRenderer {
	return &Pongo2ViewsRenderer{
		appConf:       appConf,
		logger:        logger,
		templateCache: make(map[string]*pongo2.Template),
	}
}

func (pr *Pongo2ViewsRenderer) Render(
	templateFile string,
	context map[string]any,
) ([]byte, error) {
	var pongoTemplate *pongo2.Template
	var err error = nil

	if cached, ok := pr.templateCache[templateFile]; ok {
		pongoTemplate = cached
	} else {
		pongoTemplate, err = pongo2.FromFile(templateFile)
		if err != nil {
			return []byte{}, err
		}

		if pr.appConf.RunMode == conf.RUN_MODE_RELEASE {
			pr.logger.Info(
				"Templ caching",
				templateFile,
			)
			pr.templateCache[templateFile] = pongoTemplate
		}
	}

	return pongoTemplate.ExecuteBytes(context)
}
