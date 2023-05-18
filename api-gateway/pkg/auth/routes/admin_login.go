package routes

import (
	"context"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminLoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AdminLogin(ctx *gin.Context, c pb.AuthServiceClient) {
	b := AdminLoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AdminLogin(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
