package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"

	pb "github.com/denderello/ping-pong-grpc/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	clientHost string
	clientPort string
)

func init() {
	clientCommand.Flags().StringVar(&clientHost, "clientHost", "localhost", "Host to connect to")
	clientCommand.Flags().StringVar(&clientPort, "clientPort", "8080", "Port to connect to")
}

var clientCommand = &cobra.Command{
	Use:   "client",
	Short: "Run pingpong in client mode",
	Long:  `Run pingpong in client mode and send ping messages to a pingpong server`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set up a connection to the server.
		conn, err := grpc.Dial(net.JoinHostPort(clientHost, clientPort), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)

		r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "foo"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)
	},
}
