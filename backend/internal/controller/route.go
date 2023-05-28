package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/services"
)

type Handler struct {
	userService           *services.UsersService
	articlesService       *services.ArticleService
	searchArticlesService *services.SearchArticlesService
	imageLinkService      *services.ImageLinksService
	tagService            *services.TagService
}

func NewRouter(g *gin.Engine, userService *services.UsersService, imageLinkService *services.ImageLinksService, articlesService *services.ArticleService, searchArticlesService *services.SearchArticlesService, tagService *services.TagService) {
	handler := &Handler{
		userService:           userService,
		imageLinkService:      imageLinkService,
		articlesService:       articlesService,
		searchArticlesService: searchArticlesService,
		tagService:            tagService,
	}
	RegisterHandlers(g, handler)
}
