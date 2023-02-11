package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/services"
)

type Handler struct {
	services *services.UsersService
}

func NewRouter(g *gin.Engine, services *services.UsersService) {
	handler := &Handler{
		services: services,
	}
	RegisterHandlers(g, handler)
}
