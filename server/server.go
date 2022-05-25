package server

import (
	"fmt"
	"log"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	pb "github.com/kcraley/go-grpcgreeter/greeter"
)

const (
	defaultAddress = "0.0.0.0"
	defaultPort    = "8080"
)

// server respresents the application server.
type server struct {
	address    string
	grpcServer *grpc.Server
}

// Opts respresents the configuration options that can be passed
// when creating a new application server.
type Opts struct {
	Address string
	Port    string
}

// getAddress returns the server address string from the host and port.
func (o *Opts) getAddress() string {
	var host string
	var port string

	if o.Address != "" {
		host = o.Address
	} else {
		host = defaultAddress
	}
	if o.Port != "" {
		port = o.Port
	} else {
		port = defaultPort
	}

	return fmt.Sprintf("%s:%s", host, port)
}

// New creates, initializes and returns a new application server.
func New(opts *Opts) *server {
	srv := newServerWithDefaults()

	srv.address = opts.getAddress()
	return srv
}

// newServerWithDefaults
func newServerWithDefaults() *server {
	return &server{
		address: fmt.Sprintf("%s:%s", defaultAddress, defaultPort),
		grpcServer: grpc.NewServer(
			grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
			grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		),
	}
}

// ListenAndServe starts the application server and starts handling
// HTTP request with the mux router.
func (s *server) ListenAndServe() error {
	tcpListener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("unable to create tcp listener: %v", err)
	}
	defer tcpListener.Close()

	pb.RegisterGreeterServer(s.grpcServer, &greeterServer{})

	log.Printf("application starting to listen at %s", s.address)
	return s.grpcServer.Serve(tcpListener)
}
