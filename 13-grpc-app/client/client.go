package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
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
	//doClientStreaming(ctx, clientConn)

	//server streaming
	//doServerStreaming(ctx, clientConn)

	//Bidirectional streaming
	doBiDirectionalStreaming(ctx, clientConn)
}

func doBiDirectionalStreaming(ctx context.Context, client proto.AppServiceClient) {
	stream, err := client.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	requestData := []*proto.GreetRequest{
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Magesh",
				LastName:  "Kuppan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Suresh",
				LastName:  "Kannan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Ramesh",
				LastName:  "Jayaraman",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Rajesh",
				LastName:  "Pandit",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Naresh",
				LastName:  "Kumar",
			},
		},
	}

	go func() {
		for _, req := range requestData {
			stream.Send(req)
		}
		stream.CloseSend()
	}()
	/* wg := &sync.WaitGroup{}
	wg.Add(1) */
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Thats all folks!")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Greet Result:", res.GetGreetMessage())
		}
		//done <- true
		close(done)
	}()
	<-done
}

func doServerStreaming(ctx context.Context, clientConn proto.AppServiceClient) {
	//server streaming
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	stream, err := clientConn.GeneratePrime(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		if err == io.EOF {
			break
		}
		log.Println("Received : ", res.GetNo())
	}
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
