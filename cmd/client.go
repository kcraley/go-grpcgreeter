package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kcraley/go-grpcgreeter/greeter"
)

// newClientCommand creates and returns the 'client' subcommand.
func newClientCommand() *cli.Command {
	return &cli.Command{
		Name:    "client",
		Aliases: []string{"c"},
		Usage:   "runs the application in client mode to produce messages to the server.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "address",
				Value: "127.0.0.1",
				Usage: "endpoint where the appliation server is listening.",
			},
			&cli.StringFlag{
				Name:  "name",
				Value: "world",
				Usage: "name of the individual to generate the greeting for.",
			},
			&cli.StringFlag{
				Name:  "port",
				Value: "8080",
				Usage: "port where the application server is listening.",
			},
		},
		Action: newClientAction(),
	}
}

// newClientAction
func newClientAction() cli.ActionFunc {
	return func(ctxCli *cli.Context) error {
		// Setup the grpc connection dialer.
		addr := fmt.Sprintf("%s:%s", ctxCli.String("address"), ctxCli.String("port"))
		connection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return cli.Exit(fmt.Sprintf("unable to create grpc dialer: %v", err), 1)
		}
		defer connection.Close()

		// Make our call the the server
		c := pb.NewGreeterClient(connection)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		rep, err := c.SayHello(ctx, &pb.HelloRequest{Name: ctxCli.String("name")})
		if err != nil {
			return cli.Exit(fmt.Sprintf("unable to greet: %v", err), 1)
		}
		log.Printf("Greeting Message: %s", rep.GetMessage())
		return nil
	}
}
