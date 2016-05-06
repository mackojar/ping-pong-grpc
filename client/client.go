package client

import (
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/denderello/ping-pong-grpc/pingpong"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GRPCClientConfig struct {
	Host string
	Port string
}

type GRPCClient struct {
	conf GRPCClientConfig
	cc   *grpc.ClientConn
	ci   pingpong.PingPongClient
}

func NewGRPCClient(c GRPCClientConfig) (GRPCClient, error) {
	da := net.JoinHostPort(c.Host, c.Port)

	log.Debugf("Establishing connection to %s", da)
	cc, err := grpc.Dial(da, grpc.WithInsecure())
	if err != nil {
		return GRPCClient{}, errors.Wrap(err, fmt.Sprintf("Could not connect to %s", da))
	}

	return GRPCClient{
		conf: c,
		cc:   cc,
		ci:   pingpong.NewPingPongClient(cc),
	}, nil
}

func (c GRPCClient) Close() {
	c.cc.Close()
}
