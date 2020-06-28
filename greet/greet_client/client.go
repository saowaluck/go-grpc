package main

import (
	"context"
	"fmt"
	"go-grpc/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client starting..")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close()

	client := greetpb.NewGreetingServiceClient(conn)
	clientUnary(client)
	// fmt.Println("Created client %f", client)
}

func clientUnary(client greetpb.GreetingServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "pop",
		},
	}

	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Great RPC %v", err)
	}
	log.Printf("response from greet %v", res.Result)

}
