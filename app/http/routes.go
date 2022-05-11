package http

import (
	"context"
	"github.com/agambondan/web-go-blog-grpc-rest/app/controller/client"
	"github.com/agambondan/web-go-blog-grpc-rest/app/controller/user"
	"github.com/agambondan/web-go-blog-grpc-rest/app/repo"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"os"
)

func (server *ServerHttp) routesRest(mux *runtime.ServeMux, repositories *repo.Repositories, dialOption []grpc.DialOption) {
	var err error
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	//err = pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", dialOption)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	newUserController := user.NewUserController(repositories.User)
	err = pb.RegisterUserServiceHandlerServer(context.Background(), mux, newUserController)
	if err != nil {
		log.Fatalln(err)
	}
}

func (server *ServerHttp) routesGRPC(grpcServer *grpc.Server, repositories *repo.Repositories) {
	newUserController := user.NewUserController(repositories.User)
	pb.RegisterUserServiceServer(grpcServer, newUserController)
}

func (server *ServerHttp) routesRestFiber(addr string) {
	app := fiber.New()
	app.Use(cors.New())

	//api := app.Group(viper.GetString("ENDPOINT"), middleware.Oauth2Authentication)
	api := app.Group(os.Getenv("ENDPOINT"))
	api.Get("/", client.APIIndexGet)
	api.Get("/info", client.APIInfoGet)

	log.Println("Rest Fiber is running on :", addr)
	app.Listen(addr)
}
