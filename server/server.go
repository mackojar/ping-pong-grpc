package server

import (
	"fmt"
	"net"

	"github.com/denderello/ping-pong-grpc/pingpong"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GRPCServerConfig struct {
	Port string
}

type GRPCServer struct {
	conf GRPCServerConfig
	gs   *grpc.Server
}

func NewGRPCServer(c GRPCServerConfig) *GRPCServer {
	s := &GRPCServer{
		conf: c,
		gs:   grpc.NewServer(),
	}

	pingpong.RegisterPingPongServer(s.gs, s)

	return s
}

func (s GRPCServer) Start() error {
	log.Info("Starting in server mode")

	la := net.JoinHostPort("", s.conf.Port)
	lis, err := net.Listen("tcp", la)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to listen on, %s", la))
	}

	log.Infof("Listening on %s", la)
	s.gs.Serve(lis)

	return nil
}
