package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SuhailRazi/grpc-demo/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming has sarted")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Fatalf("Sending the request with names: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client steraming finished")
	if err != nil {
		log.Fatalf("error while receving %v", err)
	}
	log.Printf("%v", res.Message)
}
