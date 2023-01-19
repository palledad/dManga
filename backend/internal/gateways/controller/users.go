package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r Handler) DeleteUser(c *gin.Context, walletAddress string) {
	c.JSON(http.StatusOK, User{})
}

func (r Handler) FindUser(c *gin.Context, walletAddress string) {
	c.JSON(http.StatusOK, User{})
}

func (r Handler) PutUser(c *gin.Context, walletAddress string) {
	c.JSON(http.StatusOK, User{})
}
