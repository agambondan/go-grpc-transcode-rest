package main

import (
	"flag"
	"github.com/agambondan/web-go-blog-grpc-rest/app/config"
	httpServer "github.com/agambondan/web-go-blog-grpc-rest/app/http"
	"github.com/agambondan/web-go-blog-grpc-rest/app/http/security"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	server                 httpServer.ServerHttp
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
	security.Init()

}

func main() {
	// it shows your line code while error
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	server.Init()
	go func() {
		// init mux server
		mux := runtime.NewServeMux(runtime.WithMarshalerOption(
			runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
				Marshaler: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						Multiline:       true,
						UseProtoNames:   true,
						EmitUnpopulated: true,
					},
					UnmarshalOptions: protojson.UnmarshalOptions{
						DiscardUnknown: true,
					},
				},
			},
		))

		// run transcoding grpc to rest
		server.RunRest(mux)

		// http server
		log.Fatalln(http.ListenAndServe("localhost:8080", mux))
	}()

	addressGRPCServer := "0.0.0.0:6060"
	// make listener for tcp protocol grcp server
	listener, err := net.Listen("tcp", addressGRPCServer)
	if err != nil {
		log.Fatalln(err)
	}

	// ssl/tls
	grpcServer := grpc.NewServer(grpc.Creds(security.LoadTLSCredentialsServer()))

	//// insecure
	//grpcServer := grpc.NewServer()

	server.RunGRPC(grpcServer)

	log.Println("Server is running on :", addressGRPCServer)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
