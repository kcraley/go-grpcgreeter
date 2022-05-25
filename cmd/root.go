package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

// Execute is the main entrypoint for the command line interface.
func Execute() error {
	return (&cli.App{
		Name:  "greeter",
		Usage: "a grpc application which handles greetings.",
		Commands: []*cli.Command{
			newServerCommand(),
			newClientCommand(),
		},
	}).Run(os.Args)
}
