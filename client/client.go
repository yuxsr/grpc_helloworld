package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "grpc_helloworld/proto"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultAddr = "localhost:50051"
	defaultName = "world"
)

func StartClient(cmd *cobra.Command, args []string) {
	addr := defaultAddr
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	if len(host) > 0 && len(port) > 0 {
		addr = fmt.Sprintf("%v:%v", host, port)
	}
	// Set up a connection to the server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(args) == 1 {
		name = args[0]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
