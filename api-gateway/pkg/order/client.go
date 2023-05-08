package order

import (
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/config"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	// using WithInsecure() because of no SSL running
	cc, err := grpc.Dial(c.OrderSuvUrl, grpc.WithInsecure()) // cc is a pointer to client connection

	if err != nil {
		fmt.Println("Could not connect : ", err)
	}
	return pb.NewOrderServiceClient(cc)
}
