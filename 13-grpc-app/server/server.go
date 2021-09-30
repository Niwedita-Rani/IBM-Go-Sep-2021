package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func (s *server) Average(stream proto.AppService_AverageServer) error {
	var sum int64
	var count int64
	for {
		req, err := stream.Recv()
		no := req.GetNo()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		log.Println("Received : ", no)
		sum += no
		count++
	}
	result := sum / count
	log.Println("Sending response : ", result)
	return stream.SendAndClose(&proto.AverageResponse{
		Result: result,
	})
}

func (s *server) GeneratePrime(req *proto.PrimeRequest, stream proto.AppService_GeneratePrimeServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	log.Println("Received Request", start, end)
	for i := start; i <= end; i++ {
		if isPrime(i) {
			time.Sleep(time.Millisecond * 500)
			log.Println("Sending Prime : ", i)
			stream.Send(&proto.PrimeResponse{
				No: i,
			})
		}
	}
	return nil
}

func (s *server) Greet(stream proto.AppService_GreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		greeting := req.GetGreeting()
		first_name := greeting.GetFirstName()
		last_name := greeting.GetLastName()
		greetMsg := fmt.Sprintf("Hi %s %s, Have a nice day!", first_name, last_name)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Sending : ", greetMsg)
		resp := &proto.GreetResponse{GreetMessage: greetMsg}
		stream.Send(resp)
	}
}

func isPrime(no int64) bool {
	if no <= 1 {
		return false
	}
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Println("Server up and running!")
}
