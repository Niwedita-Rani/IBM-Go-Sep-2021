package main

import (
	"context"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	clientConn := proto.NewAppServiceClient(client)

	//request & response
	ctx := context.Background()
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := clientConn.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.GetResult())
}
