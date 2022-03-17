package http

import (
	"google.golang.org/grpc"
)

type ServerHttp struct {
	Server *grpc.Server
}

func (server *ServerHttp) Run(grpcServer *grpc.Server) {
	server.Server = grpcServer
}
