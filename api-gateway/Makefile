proto:
# proto target compiles the protocol buffer files for authentication, order, and product packages.
# It uses the 'protoc' command to generate gRPC code from the .proto files.
# The generated code will be placed in the same directory as the .proto files.
# Compiling auth.proto
proto:
	protoc --go_out=. --go-grpc_out=. pkg/auth/pb/*.proto
	protoc --go_out=. --go-grpc_out=. pkg/order/pb/*.proto
	protoc --go_out=. --go-grpc_out=. pkg/product/pb/*.proto

server:
# server target runs the main.go file to start the server.
# Starting the server
	go run cmd/main.go