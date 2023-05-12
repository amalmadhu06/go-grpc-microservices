package services

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/product-svc/pkg/db"
	"github.com/amalmadhu06/go-grpc-microservices/product-svc/pkg/models"
	"github.com/amalmadhu06/go-grpc-microservices/product-svc/pkg/pb"
	"net/http"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	fmt.Println("Product Service : CreateProduct")
	var product models.Product
	fmt.Println("repository")
	fmt.Println(req)
	fmt.Println("----------------")
	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) FindAll(ctx context.Context, req *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	var products []models.Product

	rows, err := s.H.DB.Model(&models.Product{}).Where("stock > 0").Rows()
	defer rows.Close()
	if err != nil {
		return &pb.FindAllResponse{}, err
	}

	for rows.Next() {
		var product models.Product
		err := s.H.DB.ScanRows(rows, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	var outProducts []*pb.FindOneData
	for _, v := range products {
		var p pb.FindOneData
		p.Id = v.Id
		p.Name = v.Name
		p.Price = v.Price
		p.Stock = v.Stock

		outProducts = append(outProducts, &p)
	}

	return &pb.FindAllResponse{
		Status:   http.StatusOK,
		Products: outProducts,
	}, nil
}

func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	fmt.Println("Product Service :  DecreaseStock")
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - 1

	s.H.DB.Save(&product)

	log.OrderId = req.OrderId
	log.ProductRefer = product.Id

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
