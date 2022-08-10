package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	. "rpc-gateway/src/api/service"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
