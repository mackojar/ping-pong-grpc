package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

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
		s := server.NewGRPCServer(server.GRPCServerConfig{
			Port: serverPort,
		})

		err := s.Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}
