package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/configs"
	"github.com/palledad/dManga/backend/internal/controller"
	"github.com/palledad/dManga/backend/internal/models"
	"github.com/palledad/dManga/backend/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
)

func main() {
	r := gin.Default()

	// validatorの設定
	swagger, err := controller.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Connect to DB
	db, err := gorm.Open(postgres.Open(configs.DataSourceName))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect with DB: %v", err)
		os.Exit(1)
	}
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()

	// Create AWS Clients
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load AWS config: %v", err)
		os.Exit(1)
	}

	s3Client := s3.NewFromConfig(cfg)

	r.Use(middleware.OapiRequestValidator(swagger))
	userModel := models.NewUserModel()
	userService := services.NewUsersService(db, userModel)
	imageLinkService := services.NewImageLinksService(s3Client)
	controller.NewRouter(r, userService, imageLinkService)
	_ = r.Run()
}
