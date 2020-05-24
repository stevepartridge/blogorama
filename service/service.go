package service

import (
	"context"

	"github.com/stevepartridge/blogorama/pkg/fauxauth"
	pb "github.com/stevepartridge/blogorama/protos/go"
	_service "github.com/stevepartridge/service"
)

var (
	Name    = "blogorama"
	Version = "dev"
	Build   = "0"
	BuiltAt string
	GitHash string
)

// Service holds the server and faux auth mechanism
type Service struct {
	FauxAuth fauxauth.FauxAuth

	Server *_service.Service
}

// New creates and prepares the Service for use
func New(port int) (*Service, error) {

	err := setupDatastore()
	if err != nil {
		return nil, err
	}

	s := Service{}
	s.Server, err = _service.New(port)
	if err != nil {
		return nil, err
	}

	s.CORS()

	s.TLS()

	setWhitelist()

	err = s.fauxAuth()
	if err != nil {
		return nil, err
	}

	s.middleware()

	pb.RegisterBlogServer(s.Server.GrpcServer(), &blogService{})

	err = s.Server.AddGatewayHandler(pb.RegisterBlogHandlerFromEndpoint)
	if err != nil {
		return nil, err
	}

	s.httpMiddleware() // Add http/1 only middleware after we've established the grpc-gateway

	s.swagger()

	return &s, nil
}

// Serve handles all the heavy lifting by invoking the underlying service's serve method
func (s *Service) Serve() error {
	if err := s.Server.Serve(); err != nil {
		return err
	}
	return nil
}

// Reload manages any a hot reload of the service, typically used for config changes
func (s *Service) Reload() error {

	setupDatastore()

	return nil
}

// Shutdown is the last call of the service that sends everyone home and attempts to gracefully close down the service
func (s *Service) Shutdown(ctx context.Context) error {
	s.Server.Grpc.Server.GracefulStop()
	return nil
}
