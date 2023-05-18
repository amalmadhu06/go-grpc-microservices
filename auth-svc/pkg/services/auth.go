package services

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/db"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/models"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/pb"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/utils"
	"net/http"
)

type Server struct {
	H                                 db.Handler       // handler
	Jwt                               utils.JWTWrapper // jwt wrapper
	pb.UnimplementedAuthServiceServer                  //need to embed this for implementing
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println("Auth Service :  Register")
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("Auth Service :  Login")
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user, "user")

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	fmt.Println("Auth Service :  Validate")
	claims, err := s.Jwt.ValidateToken(req.Token, req.Role)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User
	fmt.Println(req.Role)
	if req.Role == "user" {
		if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
			return &pb.ValidateResponse{
				Status: http.StatusNotFound,
				Error:  "User not found",
			}, nil
		}
	}
	if req.Role == "admin" {
		sql := `SELECT * FROM admins WHERE email = $1 LIMIT 1;`
		if result := s.H.DB.Raw(sql, claims.Email).Scan(&user); result.Error != nil {
			return &pb.ValidateResponse{
				Status: http.StatusNotFound,
				Error:  "admin not found",
			}, nil
		}
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
