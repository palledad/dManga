package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/services"
)

type Handler struct {
	userService      *services.UsersService
	articlesService  *services.ArticleService
	imageLinkService *services.ImageLinksService
	tagService       *services.TagService
}

func NewRouter(g *gin.Engine, userService *services.UsersService, imageLinkService *services.ImageLinksService, articlesService *services.ArticleService, tagService *services.TagService) {
	handler := &Handler{
		userService:      userService,
		imageLinkService: imageLinkService,
		articlesService:  articlesService,
		tagService:       tagService,
	}
	RegisterHandlers(g, handler)
}
