run :
	go run main.go -environment development

server:
	make clean-cert
	make cert-server-side
	make clean
	make create
	make run

client:
	go run ./grpc/client/client.go

create:
	protoc --proto_path=grpc grpc/proto/*.proto --go_out=grpc/gen/
	protoc --proto_path=grpc grpc/proto/*.proto --go-grpc_out=grpc/gen/
	protoc --proto_path=grpc grpc/proto/*.proto --grpc-gateway_out=grpc/gen \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true

clean:
	rm grpc/gen/proto/*.go

cert-server-side:
	# -nodes function to skip a passphrase
	# 1. Generate CA's private key and self-signed certificate
	#openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ./app/http/security/cert/ca-key.pem -out ./app/http/security/cert/ca-cert.pem -subj "/C=ID/ST=JKT/L=JKT/O=GB, Inc./emailAddress=agamwork28@gmail.com"
	openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ./app/http/security/cert/ca-key.pem -out ./app/http/security/cert/ca-cert.pem -subj "/C=FR/ST=Occitanie/L=Toulouse/O=Tech School/OU=Education/CN=*.techschool.guru/emailAddress=techschool.guru@gmail.com"
	#openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ./app/http/security/cert/ca-key.pem -out ./app/http/security/cert/ca-cert.pem -subj "/C=ID/ST=Jakarta/L=Jakarta/O=/OU=/CN=*.go-blog/emailAddress=agamwork28@gmail.com"

	echo "CA's self-signed certificate"
	openssl x509 -in ./app/http/security/cert/ca-cert.pem -noout -text

	# 2. Generate web server's private key and certificate signing request (CSR)
	#openssl req -newkey rsa:4096 -nodes -keyout ./app/http/security/cert/server-key.pem -out ./app/http/security/cert/server-req.pem -subj "/C=ID/ST=JKT/L=JKT/O=GB, Inc./emailAddress=agamwork28@gmail.com"
	openssl req -newkey rsa:4096 -nodes -keyout ./app/http/security/cert/server-key.pem -out ./app/http/security/cert/server-req.pem -subj "/C=FR/ST=Ile de France/L=Paris/O=PC Book/OU=Computer/CN=*.pcbook.com/emailAddress=pcbook@gmail.com"
	#openssl req -newkey rsa:4096 -nodes -keyout ./app/http/security/cert/server-key.pem -out ./app/http/security/cert/server-req.pem -subj "/C=ID/ST=Jakarta/L=Jakarta/O=/OU=/CN=*.go-blog/emailAddress=agamwork28@gmail.com"

	# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
	openssl x509 -req -in ./app/http/security/cert/server-req.pem -days 365 -CA ./app/http/security/cert/ca-cert.pem -CAkey ./app/http/security/cert/ca-key.pem -CAcreateserial -out ./app/http/security/cert/server-cert.pem -extfile ./app/http/security/cert/server-ext.cnf

	echo "Server's signed certificate"
	openssl x509 -in ./app/http/security/cert/server-cert.pem -noout -text

cert-client-side:
	# 4. Generate client's private key and certificate signing request (CSR)
	openssl req -newkey rsa:4096 -nodes -keyout ./app/http/security/cert/client-key.pem -out ./app/http/security/cert/client-req.pem -subj "/C=FR/ST=Alsace/L=Strasbourg/O=PC Client/OU=Computer/CN=*.pcclient.com/emailAddress=pcclient@gmail.com"

    # 5. Use CA's private key to sign client's CSR and get back the signed certificate
	openssl x509 -req -in ./app/http/security/cert/client-req.pem -days 60 -CA ./app/http/security/cert/ca-cert.pem -CAkey ./app/http/security/cert/ca-key.pem -CAcreateserial -out ./app/http/security/cert/client-cert.pem extfile ./app/http/security/cert/client-ext.cnf

	echo "Client's signed certificate"
	openssl x509 -in ./app/http/security/cert/client-cert.pem -noout -text

clean-cert:
	rm ./app/http/security/cert/*.pem
	rm ./app/http/security/cert/*.srl

# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/protobuf/date.proto
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/protobuf/datetime.proto
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/anotations.proto