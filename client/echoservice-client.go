package main

import (
	"flag"
	"fmt"
	"log"

	"go.kimpton.io/echo/common"
	pb "go.kimpton.io/echo/echoservice"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var(
	conn *grpc.ClientConn
)

func main() {
	flag.Parse()
	
	conn, err := grpc.Dial(fmt.Sprintf(":%d", common.Port()), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	echoClient := pb.NewEchoServiceClient(conn)
	ctx := context.Background()
	pingRequest := &pb.PingRequest{
		Number: 74, 
		Message: "Hello From Client",
	}

	resp, err := echoClient.Ping(ctx, pingRequest)
	if err != nil {
		log.Fatalf("Error on Ping: %s", err)
	}

	log.Printf("Response: %s", resp.Response)
}