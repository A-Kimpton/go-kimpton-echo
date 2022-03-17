package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"go.kimpton.io/echo/common"
	pb "go.kimpton.io/echo/echoservice"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	conn *grpc.ClientConn

	endpoint = flag.String("endpoint", "", "endpoint to communicate to - blank=localhost")
	repeat   = flag.Bool("repeat", false, "if set to true, the application will continue to ping the server")

	echoClient pb.EchoServiceClient
)

func main() {
	flag.Parse()
	var err error
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", *endpoint, common.Port()), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	echoClient = pb.NewEchoServiceClient(conn)
	ctx := context.Background()
	pingRequest := &pb.PingRequest{
		Number:  74,
		Message: "Hello From Client",
	}

	var resp *pb.PingResponse
	for {

		resp, err = echoClient.Ping(ctx, pingRequest)
		if err == nil {
			break
		}
		log.Printf("Error on Ping: %s", err)
		time.Sleep(1 * time.Second)

	}

	log.Printf("Response: %s", resp.Response)

	if *repeat {
		DoRepeat(ctx)
	}

}

func DoRepeat(ctx context.Context) {

	num := uint64(0)

	// repeat..
	for {

		// build request
		pingRequest := &pb.PingRequest{
			Number:  num,
			Message: "This is a repeating request",
		}

		var resp *pb.PingResponse
		var err error
		for {

			resp, err = echoClient.Ping(ctx, pingRequest)
			if err == nil {
				break
			}
			log.Printf("Error on Ping: %s", err)
			time.Sleep(1 * time.Second)

		}

		// print resp
		log.Printf("Response: %s", resp.Response)

		// increment
		num = num + 1

		// wait 800ms
		time.Sleep(800 * time.Millisecond)

	}

}
