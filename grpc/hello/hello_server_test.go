package hello

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

const (
	port = ":1234"
)

func sayHello(con context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Receive: %v", in.GetData())
	return &HelloReply{Data: "Hello " + in.GetData()}, nil
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterHelloService(s, &HelloService{SayHello: sayHello})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
