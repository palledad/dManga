package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/gateways/repository"
)

type Handler struct {
	repository *repository.Repository
}

func NewRouter(g *gin.Engine, repo *repository.Repository) {
	handler := &Handler{
		repository: repo,
	}
	RegisterHandlers(g, handler)
}
