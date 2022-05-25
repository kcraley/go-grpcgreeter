package main

import (
	"log"

	"github.com/kcraley/go-grpcgreeter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("failed executing command: %v", err)
	}
}
