package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/backend/internal/models"
)

func (h *Handler) PostArticle(c *gin.Context) {
	articleReq := &PostArticleEntry{}
	err := c.BindJSON(&articleReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	h.articlesService.CreateArticle(&models.Article{
		Title:         articleReq.Title,
		Content:       articleReq.Content,
		Alias:         *articleReq.Alias,
		AuthorAddress: articleReq.AuthorAddress,
	})
	c.JSON(http.StatusOK, &Article{
		Id: "hoge",
	})
}

func (h *Handler) GetArticle(c *gin.Context, articleId string) {
	c.JSON(http.StatusOK, &Article{
		Id: articleId,
	})
}

func (h *Handler) UpdateArticle(c *gin.Context, articleId string) {
	c.JSON(http.StatusOK, &Article{
		Id: articleId,
	})
}

func (h *Handler) DeleteArticle(c *gin.Context, articleId string) {
	c.JSON(http.StatusOK, &Article{
		Id: articleId,
	})
}
