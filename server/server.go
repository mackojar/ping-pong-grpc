package server

import (
	"log"
	"net"

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
	r(s.server)
}

func (s GRPCServer) Start() {
	lis, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.server.Serve(lis)
}
