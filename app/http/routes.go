package http

import (
	"github.com/agambondan/web-go-blog-grpc-rest/app/controller/user"
	"github.com/agambondan/web-go-blog-grpc-rest/app/repo"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"google.golang.org/grpc"
)

func (server *ServerHttp) routes(grpcServer *grpc.Server, repositories *repo.Repositories) {
	newUserController := user.NewUserController(repositories.User)
	pb.RegisterUserServiceServer(grpcServer, newUserController)
}
