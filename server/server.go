package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	muxRouter  *mux.Router
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
		muxRouter: mux.NewRouter(),
	}
}

// Initialize registers all necessary routes and handlers that are going
// to be served by the application server.
func (s *server) Initialize() {
	pb.RegisterGreeterServer(s.grpcServer, &greeterServer{})
	grpc_prometheus.Register(s.grpcServer)

	s.muxRouter.Handle("/metrics", promhttp.Handler())
}

// ListenAndServe starts the application server and starts handling
// HTTP request with the mux router.
func (s *server) ListenAndServe() error {
	tcpListener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("unable to create tcp listener: %v", err)
	}
	defer tcpListener.Close()

	s.Initialize()

	go func() {
		httpSrv := http.Server{
			Addr:         "0.0.0.0:9092",
			Handler:      s.muxRouter,
			IdleTimeout:  10 * time.Second,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 20 * time.Second,
		}
		log.Print("serving instrumentation at /metrics")
		httpSrv.ListenAndServe()
	}()
	log.Printf("application starting to listen at %s", s.address)
	return s.grpcServer.Serve(tcpListener)
}
