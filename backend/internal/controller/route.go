package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/services"
)

type Handler struct {
	userService      *services.UsersService
	imageLinkService *services.ImageLinksService
}

func NewRouter(g *gin.Engine, userService *services.UsersService, imageLinkService *services.ImageLinksService) {
	handler := &Handler{
		userService:      userService,
		imageLinkService: imageLinkService,
	}
	RegisterHandlers(g, handler)
}
