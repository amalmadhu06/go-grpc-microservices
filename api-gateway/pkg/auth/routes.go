package auth

import (
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/auth/routes"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	fmt.Println("API Gateway :  Register Routes")
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	user := r.Group("/auth")
	user.POST("/register", svc.Register)
	user.POST("/login", svc.Login)

	admin := r.Group("/admin")
	admin.POST("/login", svc.AdminLogin)
	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}
func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) AdminLogin(ctx *gin.Context) {
	routes.AdminLogin(ctx, svc.Client)
}
