package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/denderello/ping-pong-grpc/net"
	"github.com/denderello/ping-pong-grpc/server"
)

var (
	serverPort string
        serverMsg string
)

func init() {
	serverCommand.Flags().StringVar(&serverPort, "port", "8080", "Port to listen on connections.")
	serverCommand.Flags().StringVar(&serverMsg, "msg", "pong", "Message returned by server.")
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run pingpong in server mode",
	Long:  `Run pingpong in server mode and wait for ping message to respond with a pong.`,
	Run: func(cmd *cobra.Command, args []string) {
		l := newLogger()

		s := server.NewGRPCServer(server.GRPCServerConfig{
			Logger: l,
			Address: net.NetAddress{
				Port: serverPort,
			},
			Message: serverMsg,
		})

		sigs := make(chan os.Signal, 1)
		go func() {
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			l.Debugf("Received signal %s", <-sigs)

			s.Stop()
		}()

		err := s.Start()
		if err != nil {
			l.Fatal(err)
		}
	},
}
