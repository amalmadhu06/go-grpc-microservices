package auth

import (
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/auth/pb"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	fmt.Println("API Gateway :  InitServiceClient")
	//	using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSuvUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewAuthServiceClient(cc)
}
