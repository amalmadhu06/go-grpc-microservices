package routes

import (
	"context"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/order/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductID: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
