package main

import (
	"log"
	"os"

	"github.com/denderello/ping-pong-grpc/server"
	"github.com/denderello/ping-pong-grpc/service"

	pb "github.com/denderello/ping-pong-grpc/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = ":50051"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("Please provide a server type (server, client).")
	}

	t := os.Args[1]
	switch t {
	case "server":
		runServer()
	case "client":
		runClient()
	default:
		log.Fatalf("type %s is not supported.", t)
	}
}

func runServer() {
	s := server.NewGRPCServer(server.GRPCServerConfig{
		Port: port,
	})

	s.RegisterServices(func(s *grpc.Server) {
		pb.RegisterGreeterServer(s, &service.GreetingService{})
	})

	s.Start()
}

func runClient() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(host+port, grpc.WithInsecure())
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
}
