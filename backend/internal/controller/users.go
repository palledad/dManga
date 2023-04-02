package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/models"
)

func convertToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}

	return sql.NullString{String: "", Valid: false}
}

func convertToStringPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	} else {
		return nil
	}
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

	respBody := &User{
		Biography:     convertToStringPtr(user.Biography),
		Name:          convertToStringPtr(user.Name),
		Picture:       convertToStringPtr(user.IconUrl),
		WalletAddress: user.WalletAddress,
	}
	c.JSON(http.StatusOK, respBody)
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
		Biography:     convertToNullString(user.Biography),
		Name:          convertToNullString(user.Name),
		IconUrl:       convertToNullString(user.Picture),
	})

	c.JSON(http.StatusOK, User{
		WalletAddress: walletAddress,
	})
}
