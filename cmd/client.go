package cmd

import (
	"time"

	"github.com/denderello/ping-pong-grpc/client"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	clientHost               string
	clientPort               string
	clientCycleMode          bool
	clientCycleSleepDuration time.Duration
)

func init() {
	clientCommand.Flags().StringVar(&clientHost, "host", "localhost", "Host to connect to")
	clientCommand.Flags().StringVar(&clientPort, "port", "8080", "Port to connect to")
	clientCommand.Flags().BoolVar(&clientCycleMode, "cycle-mode", false, "Wether ping requests should be automatically repeated until the client get's interrupted")
	clientCommand.Flags().DurationVar(&clientCycleSleepDuration, "cycle-sleep-duration", time.Duration(500)*time.Millisecond, "The time to wait between two consecutive ping messsages")
}

var clientCommand = &cobra.Command{
	Use:   "client",
	Short: "Run pingpong in client mode",
	Long:  `Run pingpong in client mode and send ping messages to a pingpong server`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Starting in client mode")

		c := client.NewGRPCClient(client.GRPCClientConfig{
			Host: clientHost,
			Port: clientPort,
		})
		defer c.Close()

		c.Ping(clientCycleMode, clientCycleSleepDuration)
	},
}
