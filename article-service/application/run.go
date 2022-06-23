package application

import (
	"article-service/config"
	"errors"
)

func (app *Application) Run(cfg *config.Config) error {
	return runApp(cfg, app)
}

func runApp(cfg *config.Config, app *Application) error {
	switch app.ServiceMode {
	case "grpc":
		return grpcRun(cfg)(app)
	}
	return errors.New("unrecognized mode")
}
