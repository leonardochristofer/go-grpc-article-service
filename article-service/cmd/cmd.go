package cmd

import (
	"article-service/application"
	"article-service/config"
	"article-service/route"

	"github.com/urfave/cli"
)

func runCommand(cfg *config.Config) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		app, err := application.Setup(cfg, c)
		if err != nil {
			return err
		}
		route.SetupRouter(cfg, app)
		return app.Run(cfg)
	}
}

func Cli(cfg *config.Config) *cli.App {
	clientApp := cli.NewApp()
	clientApp.Name = cfg.Application.ServiceName
	clientApp.Version = cfg.Application.ServiceVersion
	clientApp.Action = runCommand(cfg)
	clientApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "mode",
			Required: false,
			Value:    "grpc",
		},
	}
	return clientApp
}
