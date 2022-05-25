package cmd

import "github.com/urfave/cli/v2"

// newClientCommand creates and returns the 'client' subcommand.
func newClientCommand() *cli.Command {
	return &cli.Command{
		Name:    "client",
		Aliases: []string{"c"},
		Usage:   "runs the application in client mode to produce messages to the server.",
		Action:  newClientAction(),
	}
}

// newClientAction
func newClientAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		return nil
	}
}
