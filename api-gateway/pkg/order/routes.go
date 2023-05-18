package order

import (
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/auth"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/config"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	fmt.Println("API Gateway :  RegisterRoutes")
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routes := r.Group("order")
	routes.Use(a.UserAuth)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	fmt.Println("API Gateway :  CreateOrder")
	routes.CreateOrder(ctx, svc.Client)
}
