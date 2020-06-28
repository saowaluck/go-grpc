package main

import (
	"context"
	"go-grpc/calculation/calculationpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect %v", err)
	}
	defer conn.Close()

	client := calculationpb.NewCalculationServiceClient(conn)

	req := &calculationpb.CalculationRequest{
		Number1: 2,
		Number2: 135,
	}

	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error to get sum from client %v", err)
	}

	log.Printf("result is %v", res.GetSum())
}
