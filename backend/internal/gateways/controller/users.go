package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/gateways/repository"
)

func safeStringDeref(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func (r Handler) DeleteUser(c *gin.Context, walletAddress string) {
	c.JSON(http.StatusOK, User{})
}

func (r Handler) FindUser(c *gin.Context, walletAddress string) {
	res := r.repository.ReadUser(walletAddress)
	c.JSON(http.StatusOK, res)
}

func (r Handler) PutUser(c *gin.Context, walletAddress string) {
	user := PutUserEntry{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{})
		return
	}

	r.repository.CreateUser(&repository.User{
		WalletAddress: walletAddress,
		Biography:     safeStringDeref(user.Biography),
		Name:          safeStringDeref(user.Name),
		IconUrl:       safeStringDeref(user.Biography),
	})

	c.JSON(http.StatusOK, User{
		WalletAddress: walletAddress,
	})
}
