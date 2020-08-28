package hello

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:1234"
)

func TestSayHello(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &HelloRequest{Data: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetData())
}
