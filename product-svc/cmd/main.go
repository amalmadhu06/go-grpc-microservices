package main

import (
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/product-svc/pkg/config"
	"github.com/amalmadhu06/go-grpc-microservices/product-svc/pkg/db"
	"github.com/amalmadhu06/go-grpc-microservices/product-svc/pkg/pb"
	"github.com/amalmadhu06/go-grpc-microservices/product-svc/pkg/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
