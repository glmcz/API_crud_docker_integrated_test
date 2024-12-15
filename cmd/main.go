package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:  "simpleCloudService",
		Usage: "for testing purpose only",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config_file",
				Aliases: []string{"c"},
				Usage:   "set config file path",
				Value:   "./config/config.yaml",
				EnvVars: []string{"CONFIG_FILE"},
			},
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "set port",
				Value:   8080,
			},
			&cli.StringFlag{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "set template dir path",
				Value:   "./static/index.html",
			},
		},

		Action: func(c *cli.Context) error {
			if err := Run(c.Context, c.String("c"), c.Int("port"), c.String("template")); err != nil {
				return err
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		println(err.Error())
	}
}
