package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r Router) DeleteUser(c *gin.Context, walletAddress string) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (r Router) FindUser(c *gin.Context, walletAddress string) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (r Router) PutUser(c *gin.Context, walletAddress string) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
