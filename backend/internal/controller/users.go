package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/models"
)

func safeStringDeref(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func (h Handler) DeleteUser(c *gin.Context, walletAddress string) {
	user, err := h.userService.DeleteUser(walletAddress)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func (h Handler) FindUser(c *gin.Context, walletAddress string) {
	user, err := h.userService.ReadUser(walletAddress)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func (h Handler) PutUser(c *gin.Context, walletAddress string) {
	user := PutUserEntry{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	h.userService.CreateUser(&models.User{
		WalletAddress: walletAddress,
		Biography:     sql.NullString{String: safeStringDeref(user.Biography)},
		Name:          sql.NullString{String: safeStringDeref(user.Name)},
		IconUrl:       sql.NullString{String: safeStringDeref(user.Biography)},
	})

	c.JSON(http.StatusOK, User{
		WalletAddress: walletAddress,
	})
}
