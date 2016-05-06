package client

import (
	"fmt"
	"time"

	"github.com/denderello/ping-pong-grpc/pingpong"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func (c GRPCClient) Ping(cycleMode bool, cycleSleepDuration time.Duration) error {
	log.Debugf("Running with cycle mode: %t and sleep duration: %s", cycleMode, cycleSleepDuration)

	for {
		req := &pingpong.Ping{Message: "ping"}

		log.Infof("Sending message to server: %s", req.Message)
		log.Debugf("Sending request to server: %#v", req)
		resp, err := c.ci.SendPing(context.Background(), req)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Did not receive a pong."))
		}
		log.Infof("Received message from server: %s", resp.Message)

		log.Debugf("Received response from server: %#v", resp)

		if !cycleMode {
			break
		}

		time.Sleep(cycleSleepDuration)
	}

	return nil
}
