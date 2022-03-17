package main

import (
	"context"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Add(ctx context.Context, user *pb.User) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (s *server) Find(ctx context.Context, id *pb.Id) (*pb.FindResponse, error) {
	findResponse := pb.FindResponse{
		Response: &pb.Response{},
		User:     &pb.User{},
	}
	log.Println(id.Id)
	if id.Id == "1" {
		findResponse.User.Id = id.Id
		findResponse.User.Money = 10000
	}
	return &findResponse, nil
}

func (s *server) Echo(ctx context.Context, req *pb.Response) (*pb.Response, error) {
	return req, nil
}

func main() {
	// it shows your line code while error
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// make listener for tcp protocol grcp server
	listener, err := net.Listen("tcp", "localhost:6060")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Server is running")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
