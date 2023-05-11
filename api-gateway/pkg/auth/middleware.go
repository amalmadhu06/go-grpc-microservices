package auth

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	fmt.Println("API Gateway :  InitAuthMiddleware")
	return AuthMiddlewareConfig{svc: svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	fmt.Println("API Gateway :  AuthRequired")
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) > 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
