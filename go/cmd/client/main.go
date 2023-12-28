package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_example/src/service"
	"log"
)

const PORT = "10001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := service.NewDemoClient(conn)
	resp, err := client.GetInfo(context.Background(), &service.Request{Id: "0"})
	if err != nil {
		log.Fatalf("Response err: %v", err)
	}
	log.Printf("Response 1: %s", resp)

	resp, err = client.GetInfo(context.Background(), &service.Request{Id: "1"})
	if err != nil {
		log.Fatalf("Respone err: %v", err)
	}
	log.Printf("Respone 2: %s", resp)
}
