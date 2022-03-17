package main

import (
	"context"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//credentials := insecure.NewCredentials()
	conn, err := grpc.Dial("localhost:6060", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		panic(err)
	}
	client := pb.NewUserServiceClient(conn)

	findResponse, err := client.FindAll(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		log.Println(err)
	}
	log.Println(findResponse, "\n\n")
	log.Println(findResponse.Users)

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
