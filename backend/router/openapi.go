// Package router provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package router

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (DELETE /users/{wallet_address})
	DeleteUser(c *gin.Context, walletAddress string)

	// (GET /users/{wallet_address})
	FindUser(c *gin.Context, walletAddress string)

	// (PUT /users/{wallet_address})
	PutUser(c *gin.Context, walletAddress string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "wallet_address" -------------
	var walletAddress string

	err = runtime.BindStyledParameter("simple", false, "wallet_address", c.Param("wallet_address"), &walletAddress)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter wallet_address: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteUser(c, walletAddress)
}

// FindUser operation middleware
func (siw *ServerInterfaceWrapper) FindUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "wallet_address" -------------
	var walletAddress string

	err = runtime.BindStyledParameter("simple", false, "wallet_address", c.Param("wallet_address"), &walletAddress)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter wallet_address: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.FindUser(c, walletAddress)
}

// PutUser operation middleware
func (siw *ServerInterfaceWrapper) PutUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "wallet_address" -------------
	var walletAddress string

	err = runtime.BindStyledParameter("simple", false, "wallet_address", c.Param("wallet_address"), &walletAddress)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter wallet_address: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PutUser(c, walletAddress)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

	errorHandler := options.ErrorHandler

	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.DELETE(options.BaseURL+"/users/:wallet_address", wrapper.DeleteUser)

	router.GET(options.BaseURL+"/users/:wallet_address", wrapper.FindUser)

	router.PUT(options.BaseURL+"/users/:wallet_address", wrapper.PutUser)

	return router
}
