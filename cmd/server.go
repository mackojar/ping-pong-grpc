package cmd

import (
	"github.com/spf13/cobra"

	"github.com/denderello/ping-pong-grpc/pingpong"
	"github.com/denderello/ping-pong-grpc/server"
	"github.com/denderello/ping-pong-grpc/service"

	"google.golang.org/grpc"
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
		s := server.NewGRPCServer(server.GRPCServerConfig{
			Port: serverPort,
		})

		s.RegisterServices(func(s *grpc.Server) {
			pingpong.RegisterPingPongServer(s, &service.PingPongService{})
		})

		s.Start()
	},
}
