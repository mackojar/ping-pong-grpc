package client

import (
	"fmt"

	"github.com/denderello/ping-pong-grpc/log"
	"github.com/denderello/ping-pong-grpc/net"
	"github.com/denderello/ping-pong-grpc/pingpong"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GRPCClientConfig struct {
	Logger  log.Logger
	Address net.Addresser
	Cert    string
}

type GRPCClient struct {
	logger         log.Logger
	address        net.Addresser
	conn           *grpc.ClientConn
	pingPongClient pingpong.PingPongClient
}

func NewGRPCClient(conf GRPCClientConfig) (GRPCClient, error) {
	conf.Logger.Info("Starting in client mode")
	conf.Logger.Debugf("Establishing connection to %s", conf.Address.Address())
	var dialOption grpc.DialOption
	if len(conf.Cert) > 0 {
		creds, err := credentials.NewClientTLSFromFile(conf.Cert, "")
		if err != nil {
			return GRPCClient{}, errors.Wrap(err, fmt.Sprintf("Could not load TLS cert %s", conf.Cert))
		}
		dialOption = grpc.WithTransportCredentials(creds)
	} else {
		dialOption = grpc.WithInsecure()
	}
	c, err := grpc.Dial(conf.Address.Address(), dialOption)
	if err != nil {
		return GRPCClient{}, errors.Wrap(err, fmt.Sprintf("Could not connect to %s", conf.Address.Address()))
	}

	return GRPCClient{
		logger:         conf.Logger,
		address:        conf.Address,
		conn:           c,
		pingPongClient: pingpong.NewPingPongClient(c),
	}, nil
}

func (c GRPCClient) Close() {
	c.conn.Close()
}
