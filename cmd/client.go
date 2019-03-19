package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/denderello/ping-pong-grpc/client"
	"github.com/denderello/ping-pong-grpc/net"

	"github.com/spf13/cobra"
)

var (
	clientHost               string
	clientPort               string
	clientCycleMode          bool
	clientCycleSleepDuration time.Duration
	tlsClientCert            string
)

func init() {
	clientCommand.Flags().StringVar(&clientHost, "host", "localhost", "Host to connect to")
	clientCommand.Flags().StringVar(&clientPort, "port", "8080", "Port to connect to")
	clientCommand.Flags().StringVar(&tlsClientCert, "cert", "", "Server CA certificate")
	clientCommand.Flags().BoolVar(&clientCycleMode, "cycle-mode", false, "Wether ping requests should be automatically repeated until the client get's interrupted")
	clientCommand.Flags().DurationVar(&clientCycleSleepDuration, "cycle-sleep-duration", time.Duration(500)*time.Millisecond, "The time to wait between two consecutive ping messsages")
}

var clientCommand = &cobra.Command{
	Use:   "client",
	Short: "Run pingpong in client mode",
	Long:  `Run pingpong in client mode and send ping messages to a pingpong server`,
	Run: func(cmd *cobra.Command, args []string) {
		l := newLogger()

		c, err := client.NewGRPCClient(client.GRPCClientConfig{
			Logger: l,
			Address: net.NetAddress{
				Host: clientHost,
				Port: clientPort,
			},
			Cert: tlsClientCert,
		})
		if err != nil {
			l.Fatal(err)
		}

		sigs := make(chan os.Signal, 1)
		go func() {
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			l.Debugf("Received signal %s", <-sigs)
			c.Close()
		}()

		defer c.Close()

		err = c.Ping(clientCycleMode, clientCycleSleepDuration)
		if err != nil {
			l.Fatal(err)
		}
	},
}
