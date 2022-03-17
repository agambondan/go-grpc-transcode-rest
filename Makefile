create:
	protoc --proto_path=grpc grpc/proto/*.proto --go_out=grpc/gen/
	protoc --proto_path=grpc grpc/proto/*.proto --go-grpc_out=grpc/gen/

clean:
	rm grpc/gen/proto/*.go