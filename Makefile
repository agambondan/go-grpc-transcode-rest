create:
	protoc --proto_path=grpc grpc/proto/*.proto --go_out=grpc/gen/
	protoc --proto_path=grpc grpc/proto/*.proto --go-grpc_out=grpc/gen/
	protoc -I . --grpc-gateway_out ./grpc/gen \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		grpc/proto/user.proto

clean:
	rm grpc/gen/proto/*.go

# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/protobuf/date.proto
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/protobuf/datetime.proto
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/anotations.proto