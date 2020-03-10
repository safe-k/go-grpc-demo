package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/safe-k/go-grpc-demo/proto"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterCalculatorServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, r *proto.Request) (*proto.Response, error) {
	a, b := r.GetA(), r.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, r *proto.Request) (*proto.Response, error) {
	a, b := r.GetA(), r.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}
