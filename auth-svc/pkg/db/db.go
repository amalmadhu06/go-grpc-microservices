package db

import (
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(
		&models.User{},
	)
	return Handler{DB: db}

}
