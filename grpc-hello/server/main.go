package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"rpctest/grpc-hello/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	proto.UnimplementedGreeterServer
}

// SayHello implements Greeter.SayHello
func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	name := req.GetName()
	if name == "" {
		name = "world"
	}
	return &proto.HelloReply{Message: fmt.Sprintf("Hello, %s!", name)}, nil
}

func main() {
	addr := flag.String("addr", ":50051", "server listen address")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGreeterServer(grpcServer, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Printf("gRPC server listening on %s", *addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
