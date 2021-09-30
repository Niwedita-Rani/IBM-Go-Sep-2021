package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"

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
