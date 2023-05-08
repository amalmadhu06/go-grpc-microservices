proto:
# proto target compiles the protocol buffer files for authentication, order, and product packages.
# It uses the 'protoc' command to generate gRPC code from the .proto files.
# The generated code will be placed in the same directory as the .proto files.
# Compiling auth.proto
	protoc pkg/auth/pb/auth.proto --go-grpc_out=.

# Compiling order.proto
	protoc pkg/order/pb/order.proto --go-grpc_out=.
# Compiling product.proto
	protoc pkg/product/pb/product.proto --go-grpc_out=.

server:
# server target runs the main.go file to start the server.
# Starting the server
	go run cmd/main.go