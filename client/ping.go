package client

import (
	"fmt"
	"time"

	"github.com/denderello/ping-pong-grpc/pingpong"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func (c GRPCClient) Ping(cycleMode bool, cycleSleepDuration time.Duration) error {
	c.logger.Debugf("Running with cycle mode: %t and sleep duration: %s", cycleMode, cycleSleepDuration)

	for {
		req := &pingpong.Ping{Message: "ping"}

		c.logger.Infof("Sending message to server: %s", req.Message)
		c.logger.Debugf("Sending request to server: %#v", req)
		resp, err := c.pingPongClient.SendPing(context.Background(), req)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Did not receive a pong."))
		}
		c.logger.Infof("Received message from server: %s", resp.Message)

		c.logger.Debugf("Received response from server: %#v", resp)

		if !cycleMode {
			break
		}

		time.Sleep(cycleSleepDuration)
	}

	return nil
}
