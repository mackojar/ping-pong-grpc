package server

import (
	"fmt"
	"net"

	lnet "github.com/denderello/ping-pong-grpc/net"
	"github.com/denderello/ping-pong-grpc/pingpong"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	a  lnet.Addresser
	gs *grpc.Server
}

func NewGRPCServer(a lnet.Addresser) *GRPCServer {
	s := &GRPCServer{
		a:  a,
		gs: grpc.NewServer(),
	}

	pingpong.RegisterPingPongServer(s.gs, s)

	return s
}

func (s GRPCServer) Start() error {
	log.Info("Starting in server mode")

	lis, err := net.Listen("tcp", s.a.Address())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to listen on, %s", s.a.Address()))
	}

	log.Infof("Listening on %s", s.a.Address())
	s.gs.Serve(lis)

	return nil
}

func (s GRPCServer) Stop() {
	log.Info("Stopping server")
	s.gs.Stop()
}
