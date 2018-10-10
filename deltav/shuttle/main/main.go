package main

import (
	"log"
	"time"

	protos "github.com/golang001/deltav/protos"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewWorldModelClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Register(ctx, &protos.RegisterRequest{})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}
	log.Printf("Greeting: %s", r.Effect)
}
