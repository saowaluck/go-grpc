package main

import (
	"context"
	"go-grpc/calculation/calculationpb"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculationpb.CalculationRequest) (*calculationpb.CalculationResponse, error) {
	log.Printf("request from client %v\n", req)
	sum := req.Number1 + req.Number2

	return &calculationpb.CalculationResponse{
		Sum: sum,
	}, nil
}

func main() {
	log.Println("start calculation server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error when listen %v", err)
	}

	s := grpc.NewServer()
	calculationpb.RegisterCalculationServiceServer(s, &server{})

	err = s.Serve(lis)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve %v", err)
	}
}
