package cmd

import (
	"log"

	"github.com/kcraley/go-grpcgreeter/server"
	"github.com/urfave/cli/v2"
)

// newServerCommand creates and returns the 'server' subcommand.
func newServerCommand() *cli.Command {
	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "starts the grpc server to handle incoming requests from clients.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "port",
				Value: "8080",
				Usage: "port the application server is listening on.",
			},
		},
		Action: newServerAction(),
	}
}

// newServerAction handles the main logic of the 'server' subcommand.
func newServerAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		srv := server.New(&server.Opts{})

		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("failed running application server: %v", err)
			return err
		}

		return nil
	}
}
