package views

import (
	"github.com/Mth-Ryan/waveaction/pkg/logger"
	"github.com/flosch/pongo2/v6"
)

type Pongo2ViewsRenderer struct {
	logger        logger.ApplicationLogger
	templateCache map[string](*pongo2.Template)
}

func NewPongo2ViewsRenderer(logger logger.ApplicationLogger) *Pongo2ViewsRenderer {
	return &Pongo2ViewsRenderer{
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
		
		pr.logger.Info(
			"Templ caching",
			templateFile,
		)
		pr.templateCache[templateFile] = pongoTemplate
	}

	return pongoTemplate.ExecuteBytes(context)
}
