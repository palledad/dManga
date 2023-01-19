package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/gateways/controller"

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
	r.Use(middleware.OapiRequestValidator(swagger))
	controller.NewRouter(r)
	_ = r.Run()
}
