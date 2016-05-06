package server

import (
	"net"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCServerConfig struct {
	Port string
}

type GRPCServer struct {
	config GRPCServerConfig
	server *grpc.Server
}

type GRPCServiceRegistrator func(*grpc.Server)

func NewGRPCServer(c GRPCServerConfig) GRPCServer {
	return GRPCServer{
		config: c,
		server: grpc.NewServer(),
	}
}

func (s GRPCServer) RegisterServices(r GRPCServiceRegistrator) {
	log.Debug("Registering gRPC services for server.")

	r(s.server)
}

func (s GRPCServer) Start() {
	log.Info("Starting in server mode")

	la := net.JoinHostPort("", s.config.Port)
	lis, err := net.Listen("tcp", la)
	if err != nil {
		log.Fatalf("Failed to listen on, %s with error: %v", la, err)
	}

	log.Infof("Listening on %s", la)
	s.server.Serve(lis)
}
