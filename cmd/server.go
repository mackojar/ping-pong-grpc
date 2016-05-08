package cmd

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/denderello/ping-pong-grpc/net"
	"github.com/denderello/ping-pong-grpc/server"
)

var (
	serverPort string
)

func init() {
	serverCommand.Flags().StringVar(&serverPort, "port", "8080", "Port to listen on connections.")
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run pingpong in server mode",
	Long:  `Run pingpong in server mode and wait for ping message to respond with a pong.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewGRPCServer(net.NetAddress{
			Port: serverPort,
		})

		sigs := make(chan os.Signal, 1)
		go func() {
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			log.Debugf("Received signal %s", <-sigs)

			s.Stop()
		}()

		err := s.Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}
