package server

import (
	"fmt"
	"net"

	"github.com/denderello/ping-pong-grpc/log"
	lnet "github.com/denderello/ping-pong-grpc/net"
	"github.com/denderello/ping-pong-grpc/pingpong"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GRPCServerConfig struct {
	Logger  log.Logger
	Address lnet.Addresser
	Message string
	Cert    string
	Key     string
}

type GRPCServer struct {
	logger  log.Logger
	address lnet.Addresser
	server  *grpc.Server
	message string
}

func NewGRPCServer(c GRPCServerConfig) *GRPCServer {
	var gs *grpc.Server
	if len(c.Key) > 0 {
		creds, err := credentials.NewServerTLSFromFile(c.Cert, c.Key)
		if err != nil {
			c.Logger.Errorf("Could not load TLS keys: %s", err)
			return nil
		}
		gs = grpc.NewServer(grpc.Creds(creds))
		c.Logger.Infof("TLS mode enabled")
	} else {
		gs = grpc.NewServer()
	}
	s := &GRPCServer{
		logger:  c.Logger,
		address: c.Address,
		server:  gs,
		message: c.Message,
	}

	pingpong.RegisterPingPongServer(s.server, s)

	return s
}

func (s GRPCServer) Start() error {
	s.logger.Info("Starting server mode")

	lis, err := net.Listen("tcp", s.address.Address())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to listen on, %s", s.address.Address()))
	}

	s.logger.Infof("Listening on %s", s.address.Address())
	s.server.Serve(lis)

	return nil
}

func (s GRPCServer) Stop() {
	s.logger.Info("Stopping server")
	s.server.Stop()
}
