package db

import (
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/config"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/models"
	"github.com/amalmadhu06/go-grpc-microservices/auth-svc/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DB *gorm.DB
}

func Init(config config.Config) Handler {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(
		&models.User{},
		&models.Admin{},
	)
	// check if admin already exists in db
	var count int
	sql := `SELECT COUNT(*) FROM admins WHERE email = $1`
	if err := db.Raw(sql, config.AdminEmail).Scan(&count).Error; err != nil {
		log.Fatalln(err)
	}
	if count == 0 {
		// create admin in db using env variables
		sql := `INSERT INTO admins(email, password) VALUES ($1, $2)`
		if err := db.Exec(sql, config.AdminEmail, utils.HashPassword(config.AdminPassword)).Error; err != nil {
			log.Fatalln(err)
		}
	}
	return Handler{DB: db}

}
