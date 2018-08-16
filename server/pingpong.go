package server

import (
	"github.com/denderello/ping-pong-grpc/pingpong"

	"golang.org/x/net/context"
)

func (s *GRPCServer) SendPing(ctx context.Context, in *pingpong.Ping) (*pingpong.Pong, error) {
	return &pingpong.Pong{Message: s.message}, nil
}
