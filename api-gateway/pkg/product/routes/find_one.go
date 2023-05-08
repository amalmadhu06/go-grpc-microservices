package routes

import (
	"context"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/product/pb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FineOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
