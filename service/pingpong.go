package service

import (
	"github.com/denderello/ping-pong-grpc/pingpong"

	"golang.org/x/net/context"
)

type PingPongService struct{}

func (pps *PingPongService) SendPing(ctx context.Context, in *pingpong.Ping) (*pingpong.Pong, error) {
	return &pingpong.Pong{Message: "pong"}, nil
}
