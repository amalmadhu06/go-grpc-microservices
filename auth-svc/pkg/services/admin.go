package services

import (
	"context"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/models"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/pb"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/utils"
	"net/http"
)

func (s *Server) AdminLogin(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var admin models.Admin

	if result := s.H.DB.Where(&models.Admin{Email: req.Email}).First(&admin); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "admin not found",
		}, nil
	}

	if match := utils.CheckPasswordHash(req.Password, admin.Password); !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "invalid credentials ",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(models.User(admin), "admin")
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil

}
