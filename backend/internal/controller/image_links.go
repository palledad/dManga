package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetImageLink(c *gin.Context) {
	now := time.Now().Unix()
	now_str := strconv.FormatInt(now, 10)
	url, err := h.imageLinkService.GetLink(c, now_str)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	resp := &ImageLink{
		Link: url,
	}

	c.JSON(http.StatusOK, resp)
}
