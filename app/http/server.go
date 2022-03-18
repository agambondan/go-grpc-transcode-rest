package http

import (
	"github.com/agambondan/web-go-blog-grpc-rest/app/repo"
	"google.golang.org/grpc"
	"log"
)

type ServerHttp struct {
	Server *grpc.Server
}

func (server *ServerHttp) Run(grpcServer *grpc.Server) {
	newRepositories, err := repo.NewRepositories()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	server.Server = grpcServer
	server.routes(server.Server, newRepositories)
}
