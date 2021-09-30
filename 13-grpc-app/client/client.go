package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	clientConn := proto.NewAppServiceClient(client)
	ctx := context.Background()
	//request & response
	//doRequestResponse(ctx, clientConn)

	//client streaming
	doClientStreaming(ctx, clientConn)
}

func doClientStreaming(ctx context.Context, clientConn proto.AppServiceClient) {
	nos := []int64{3, 1, 4, 2, 5, 6, 7, 8, 9, 10}
	stream, err := clientConn.Average(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		time.Sleep(time.Millisecond * 500)
		log.Println("Sending : ", no)
		req := &proto.AverageRequest{No: no}
		err := stream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.GetResult())
}

func doRequestResponse(ctx context.Context, clientConn proto.AppServiceClient) {
	//request & response

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
