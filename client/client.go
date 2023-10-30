package main

import (
	"context"
	"log"

	"Example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	log.Printf("Connection State: %s", conn.GetState().String())
	defer conn.Close()

	ServiceConn := proto.NewExampleServiceClient(conn)

	request := new(proto.EmptyRequest)
	res, err := ServiceConn.ExampleMethod(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling ExampleMethod: %s", err)
	}
	log.Printf("Response from server: %s", res)
}
