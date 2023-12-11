package main

import (
	"context"
	"gate-service/runner"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	ctx := context.TODO()

	app := &cli.App{
		Name:  "Gate Service",
		Usage: "As a TCP Server to serve gate controller",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "Load configuration from `FILE`",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			runner.Run(ctx, cCtx.Value("config").(string))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
