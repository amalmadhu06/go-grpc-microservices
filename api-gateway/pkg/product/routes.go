package product

import (
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/auth"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/config"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	fmt.Println("API Gateway :  RegisterRoutes")
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
	routes.GET("/", svc.FindAll)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	fmt.Println("API Gateway :  FindOne")
	routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	fmt.Println("API Gateway :  CreateProduct called --> 1")
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) FindAll(ctx *gin.Context) {
	routes.FindAll(ctx, svc.Client)
}
