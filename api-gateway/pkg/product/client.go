package product

import (
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/config"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/product/pb"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProductSuvUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
