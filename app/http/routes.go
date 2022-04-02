package http

import (
	"context"
	"github.com/agambondan/web-go-blog-grpc-rest/app/controller/user"
	"github.com/agambondan/web-go-blog-grpc-rest/app/repo"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
)

func (server *ServerHttp) routesRest(mux *runtime.ServeMux, repositories *repo.Repositories, dialOption []grpc.DialOption) {
	newUserController := user.NewUserController(repositories.User)
	err := pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, "0.0.0.0:8080", dialOption)
	if err != nil {
		log.Println(err)
	}
	err = pb.RegisterUserServiceHandlerServer(context.Background(), mux, newUserController)
	if err != nil {
		log.Println(err)
	}
}

func (server *ServerHttp) routesGRPC(grpcServer *grpc.Server, repositories *repo.Repositories) {
	newUserController := user.NewUserController(repositories.User)
	pb.RegisterUserServiceServer(grpcServer, newUserController)
}
