package client

import (
	"fmt"

	"github.com/denderello/ping-pong-grpc/net"
	"github.com/denderello/ping-pong-grpc/pingpong"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	a  net.Addresser
	cc *grpc.ClientConn
	ci pingpong.PingPongClient
}

func NewGRPCClient(a net.Addresser) (GRPCClient, error) {
	log.Debugf("Establishing connection to %s", a.Address())
	cc, err := grpc.Dial(a.Address(), grpc.WithInsecure())
	if err != nil {
		return GRPCClient{}, errors.Wrap(err, fmt.Sprintf("Could not connect to %s", a.Address()))
	}

	return GRPCClient{
		a:  a,
		cc: cc,
		ci: pingpong.NewPingPongClient(cc),
	}, nil
}

func (c GRPCClient) Close() {
	c.cc.Close()
}
