package main

import (
	"context"
	"log"
	"time"

	example "grpc-client/proto" // Importe o pacote gerado pelo compilador protobuf

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	c := example.NewExampleServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.SayHello(ctx, &example.HelloRequest{Name: "John"})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	log.Println("Response:", res.Message)
}
