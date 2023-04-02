package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/services"
)

type Handler struct {
	userService      *services.UsersService
	articlesService  *services.ArticleService
	imageLinkService *services.ImageLinksService
}

func NewRouter(g *gin.Engine, userService *services.UsersService, imageLinkService *services.ImageLinksService, articlesService *services.ArticleService) {
	handler := &Handler{
		userService:      userService,
		imageLinkService: imageLinkService,
		articlesService:  articlesService,
	}
	RegisterHandlers(g, handler)
}
