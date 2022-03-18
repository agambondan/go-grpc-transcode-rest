package user

import (
	"github.com/agambondan/web-go-blog-grpc-rest/app/repo"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
)

type Controller struct {
	userRepository repo.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserController(repo repo.UserRepository) *Controller {
	return &Controller{userRepository: repo}
}
