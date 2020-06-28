package main

import (
	"context"
	"fmt"
	greetpb "go-grpc/greet/greetpb"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Great function was invoked with %v", req)
	firstname := req.GetGreeting().GetFirstName()
	result := "Hello" + firstname

	return &greetpb.GreetResponse{
		Result: result,
	}, nil
}

func main() {
	fmt.Println("Server is started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Fail to lister %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve %v", err)
	}
}
