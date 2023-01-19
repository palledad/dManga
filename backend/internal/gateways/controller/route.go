package controller

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewRouter(g *gin.Engine) {
	RegisterHandlers(g, &Handler{})
}
