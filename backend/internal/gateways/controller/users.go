package controller

import (
	"database/sql"
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
		Biography:     sql.NullString{String: safeStringDeref(user.Biography)},
		Name:          sql.NullString{String: safeStringDeref(user.Name)},
		IconUrl:       sql.NullString{String: safeStringDeref(user.Biography)},
	})

	c.JSON(http.StatusOK, User{
		WalletAddress: walletAddress,
	})
}
