package main

import (
	"flag"
	"github.com/agambondan/web-go-blog-grpc-rest/app/config"
	"github.com/agambondan/web-go-blog-grpc-rest/app/http"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	server                 http.ServerHttp
	configuration          config.Configuration
	pathFileEnvDevelopment = "./.env.development"
	pathFileEnvProduction  = "./.env.production"
)

func init() {
	env := flag.String("environment", "", "set environment")
	flag.Parse()
	switch *env {
	case "development":
		if err := godotenv.Load(pathFileEnvDevelopment); err != nil {
			log.Println("no env gotten")
		}
		os.Setenv("ENVIRONMENT", "development")
	default:
		if err := godotenv.Load(pathFileEnvProduction); err != nil {
			log.Println("no env gotten")
		}
		os.Setenv("ENVIRONMENT", "production")
	}
	configuration.Init()
	log.Println(config.Config)

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

	//grpcServer := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&security.Cert)))
	grpcServer := grpc.NewServer()

	server.Run(grpcServer)
	//pb.RegisterUserServiceServer(grpcServer)

	log.Println("Server is running")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
