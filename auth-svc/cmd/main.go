package main

import (
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/config"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/db"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/pb"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/services"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/utils"

	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed at config ", err)
	}

	fmt.Println(c.DBUrl)
	h := db.Init(c.DBUrl)

	jwt := utils.JWTWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("failed at listening : ", err)
	}
	fmt.Println("Auth svc on ", c.Port)
	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
