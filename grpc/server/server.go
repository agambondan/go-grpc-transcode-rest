package main

import (
	"context"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Add(ctx context.Context, user *pb.User) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (s *server) FindAll(ctx context.Context, emptyRequest *pb.EmptyRequest) (*pb.FindAllResponse, error) {
	var users []*pb.User
	t := time.Now()
	users = append(users, &pb.User{
		Base: &pb.BaseUUID{
			Id: uuid.New().String(),
			Time: &pb.BaseDate{
				CreatedAt: t.String()[:19],
				UpdatedAt: t.String()[:19],
				DeletedAt: t.String()[:19],
			},
		},
		FullName:    "Firman Agam",
		Gender:      "Male",
		Email:       "agamwork28@gmail.com",
		PhoneNumber: "081214025919",
		Username:    "agambondan",
		Password:    "agambondan",
	})
	findAllResponse := pb.FindAllResponse{
		Response: &pb.Response{
			Status:  true,
			Message: "Data Found",
			Error:   "",
		},
		Users: users,
	}
	return &findAllResponse, nil
}

func (s *server) Echo(ctx context.Context, req *pb.Response) (*pb.Response, error) {
	return req, nil
}

func main() {
	// it shows your line code while error
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// make listener for tcp protocol grpc server
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
