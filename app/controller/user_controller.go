package controller

import (
	"context"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
)

type UserServer struct{}

func (u *UserServer) Add(ctx context.Context, user *pb.User) (*pb.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServer) FindAll(ctx context.Context, request *pb.PaginateRequest) (*pb.FindAllResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServer) Find(ctx context.Context, request *pb.PaginateRequest) (*pb.FindResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewUserServer() pb.UserServiceServer {
	return new(UserServer)
}
