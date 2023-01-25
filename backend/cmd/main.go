package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/configs"
	"github.com/palledad/dManga/backend/internal/gateways/controller"
	"github.com/palledad/dManga/backend/internal/gateways/repository"
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

	tmp, _ := db.DB()

	tmp.Ping()

	r.Use(middleware.OapiRequestValidator(swagger))

	repository := repository.NewRepository(db)
	controller.NewRouter(r, repository)
	_ = r.Run()
}
