package main

import (
	"context"
	"flag"
	"log"
	"rpctest/grpc-hello/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "server address")
	name := flag.String("name", "world", "name to greet")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	resp, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %s", resp.GetMessage())
}
