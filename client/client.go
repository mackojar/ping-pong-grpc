package client

import (
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/denderello/ping-pong-grpc/pingpong"
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

func NewGRPCClient(c GRPCClientConfig) GRPCClient {
	da := net.JoinHostPort(c.Host, c.Port)
	log.Debugf("Establishing connection to %s", da)
	cc, err := grpc.Dial(da, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not ccect to %s with error: %v", da, err)
	}

	return GRPCClient{
		conf: c,
		cc:   cc,
		ci:   pingpong.NewPingPongClient(cc),
	}
}

func (c GRPCClient) Close() {
	c.cc.Close()
}
