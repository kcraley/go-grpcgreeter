package server

import (
	"context"
	"log"

	pb "github.com/kcraley/go-grpcgreeter/greeter"
)

// server implements the grpc application server.
type greeterServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello implementes the greeter.GreeterServer
func (s *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
