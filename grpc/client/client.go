package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/agambondan/web-go-blog-grpc-rest/app/config"
	"github.com/agambondan/web-go-blog-grpc-rest/app/http/security"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"os"
)

var (
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
	// init global variable security
	security.Init()

}

func main() {
	// it shows your line code while print
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// connection to server grpc
	//conn, err := grpc.Dial("0.0.0.0:6060", grpc.WithTransportCredentials(security.LoadTLSCredentialsClient()))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT_GRPC")), grpc.WithTransportCredentials(security.CredTransportClient))
	if err != nil {
		log.Println(err)
	}

	// make client for user service
	client := pb.NewUserServiceClient(conn)

	findResponse, err := client.FindAll(context.Background(), &pb.PaginateRequest{})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("\nStruct Value\n", findResponse.GetStructValue())
		log.Println("\nList Value\n", findResponse.GetListValue())
	}
}
