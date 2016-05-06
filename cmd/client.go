package cmd

import (
	"net"
	"time"

	pb "github.com/denderello/ping-pong-grpc/helloworld"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
		log.Debugf("Running with cycle mode: %t and sleep duration: %s", clientCycleMode, clientCycleSleepDuration)

		da := net.JoinHostPort(clientHost, clientPort)
		log.Debugf("Establishing connection to %s", da)
		conn, err := grpc.Dial(da, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect to %s with error: %v", da, err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)

		for {
			req := &pb.HelloRequest{Name: "foo"}
			log.Debugf("Sending request to server: %#v", req)
			resp, err := c.SayHello(context.Background(), req)
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Infof("Received message from server: %s", resp.Message)
			log.Debugf("Received response from server: %#v", resp)

			if !clientCycleMode {
				break
			}

			time.Sleep(clientCycleSleepDuration)
		}
	},
}
