package router

import "github.com/gin-gonic/gin"

type Router struct {
}

func NewRouter(g *gin.Engine) {
	RegisterHandlers(g, &Router{})
}
