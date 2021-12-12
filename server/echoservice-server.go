package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "go.kimpton.io/echo/echoservice"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"go.kimpton.io/echo/common"
)

type EchoServer struct {}

func main() {
	flag.Parse()

	log.Printf("Starting Echo.EchoService on port %d", common.Port())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", common.Port()))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", common.Port(), err)
	}

	es := EchoServer{}	

	grpcServer := grpc.NewServer()

	pb.RegisterEchoServiceServer(grpcServer, &es)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %d: %v", common.Port(), err)
	}
}

func (es *EchoServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received Ping Request: %s", req.Message)
	return &pb.PingResponse{
		Response: fmt.Sprintf("Message Received Num: [%d], Message: [%s]", req.Number, req.Message),
		}, nil
}