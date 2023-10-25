package views

import (
	"io"
	"path"

	"github.com/Mth-Ryan/go-web-templates/pkg/conf"
	"github.com/Mth-Ryan/go-web-templates/pkg/logger"
	"github.com/flosch/pongo2/v6"
)

type Pongo2ViewsFactory struct {
	appConf *conf.AppConf
	logger  logger.ApplicationLogger
}

func NewPongo2ViewsFactory(
	appConf *conf.AppConf,
	logger logger.ApplicationLogger,
) *Pongo2ViewsFactory {
	return &Pongo2ViewsFactory{
		appConf,
		logger,
	}
}

func (f *Pongo2ViewsFactory) GetRenderer(folder string, extension string) ViewsRenderer {
	return &Pongo2ViewsRenderer{
		templatesFolder: folder,
		templatesExtension: extension,
		useCache: f.appConf.RunMode == conf.RUN_MODE_RELEASE,
		logger: f.logger,
		templateCache: make(map[string]*pongo2.Template),
	}
}

type Pongo2ViewsRenderer struct {
	templatesFolder    string
	templatesExtension string
	useCache           bool
	logger             logger.ApplicationLogger
	templateCache      map[string](*pongo2.Template)
}

func (pr *Pongo2ViewsRenderer) Load() error {
	return nil
}

func (pr *Pongo2ViewsRenderer) Render(
	writer io.Writer,
	template string,
	context map[string]any,
) error {
	templateFullName := path.Join(
		pr.templatesFolder,
		template+pr.templatesExtension,
	)

	var pongoTemplate *pongo2.Template
	var err error = nil

	if cached, ok := pr.templateCache[templateFullName]; ok {
		pongoTemplate = cached
		return pongoTemplate.ExecuteWriter(context, writer)
	} else {
		pongoTemplate, err = pongo2.FromFile(templateFullName)
		if err != nil {
			return err
		}


		if err := pongoTemplate.ExecuteWriter(context, writer); err != nil {
			return err
		}

		if pr.useCache {
			pr.logger.Info(
				"Templ caching",
				templateFullName,
			)
			pr.templateCache[templateFullName] = pongoTemplate
		}

		return nil
	}
}

