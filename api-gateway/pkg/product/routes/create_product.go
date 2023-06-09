package routes

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	fmt.Println("API Gateway :  CreateProduct")
	var body CreateProductRequestBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println("api gateway")
	fmt.Println("body.stock : ", body.Stock)
	fmt.Println(body)
	fmt.Println("-----------------")
	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
