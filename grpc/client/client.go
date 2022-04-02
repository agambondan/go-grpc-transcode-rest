package main

import (
	"context"
	"github.com/agambondan/web-go-blog-grpc-rest/app/http/security"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// init global variable security
	security.Init()
	structpb.Struct{}

	// connection to server grpc
	conn, err := grpc.Dial("0.0.0.0:6060", grpc.WithTransportCredentials(security.LoadTLSCredentialsClient()))
	if err != nil {
		log.Println(err)
	}

	// make client for user service
	client := pb.NewUserServiceClient(conn)

	findResponse, err := client.FindAll(context.Background(), &pb.PaginateRequest{})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(findResponse.Users[0], "\n\n")
	}

	//response, err := client.Add(context.Background(), &pb.User{Id: "1", Money: 1000})
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(response)
	//log.Println(response.Message)
	//
	//findResponse, err := client.Find(context.Background(), &pb.Id{Id: "1"})
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(findResponse)
	//log.Println(findResponse.User)
	//
	//findResponse, err = client.Find(context.Background(), &pb.Id{Id: "2"})
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(findResponse)
	//log.Println(findResponse.User)
	//
	//response, err = client.Echo(context.Background(), &pb.Response{Message: "Hello World"})
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(response)
	//log.Println(response.Message)
}
