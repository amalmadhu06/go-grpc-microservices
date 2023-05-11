package services

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/order-svc/pkg/client"
	"github.com/amalmadhu06/go-grpc-microservices/order-svc/pkg/db"
	"github.com/amalmadhu06/go-grpc-microservices/order-svc/pkg/models"
	"github.com/amalmadhu06/go-grpc-microservices/order-svc/pkg/pb"
	"net/http"
)

type Server struct {
	H          db.Handler
	ProductSvc client.ProductServiceClient
	pb.UnimplementedOrderServiceServer
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	fmt.Println("Order Service :  CreateOrder")
	product, err := s.ProductSvc.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest, Error: err.Error(),
		}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: product.Status,
			Error:  product.Error,
		}, nil
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  "Stock too less",
		}, nil
	}

	order := models.Order{
		Price:     product.Data.Price,
		ProductId: product.Data.Id,
		UserId:    req.UserId,
	}

	s.H.DB.Create(&order)

	res, err := s.ProductSvc.DecreaseStock(req.ProductId, order.Id)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
