package http

import (
	"github.com/agambondan/web-go-blog-grpc-rest/app/repo"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
)

type ServerHttp struct {
	Server       *grpc.Server
	Mux          *runtime.ServeMux
	Repositories *repo.Repositories
}

func (server *ServerHttp) Init() {
	repositories, err := repo.NewRepositories()
	if err != nil {
		log.Fatalln(err)
	}
	server.Repositories = repositories
}

func (server *ServerHttp) RunRest(mux *runtime.ServeMux, dialOptions []grpc.DialOption) {
	server.Mux = mux
	server.routesRest(mux, server.Repositories, dialOptions)
}

func (server *ServerHttp) RunGRPC(grpcServer *grpc.Server) {
	server.Server = grpcServer
	server.routesGRPC(server.Server, server.Repositories)
}
