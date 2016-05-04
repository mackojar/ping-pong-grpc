package service

import (
	pb "github.com/denderello/ping-pong-grpc/helloworld"

	"golang.org/x/net/context"
)

// server is used to implement helloworld.GreeterServer.
type GreetingService struct{}

// SayHello implements helloworld.GreeterServer
func (s *GreetingService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
