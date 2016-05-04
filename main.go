package main

import (
	"log"
	"net"
	"os"

	pb "github.com/denderello/ping-pong-grpc/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
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
