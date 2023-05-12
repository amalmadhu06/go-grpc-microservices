package routes

import (
	"context"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAll(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.FindAll(context.Background(), &pb.FindAllRequest{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}
