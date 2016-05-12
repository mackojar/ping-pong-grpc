package cmd

import (
	"fmt"

	"github.com/denderello/ping-pong-grpc/log"

	"github.com/Sirupsen/logrus"
)

func printProjectVersion() {
	fmt.Printf("PingPong %s (%s)\n", projectVersion, projectCommit)
}

func newLogger() log.Logger {
	return logrus.StandardLogger()
}
