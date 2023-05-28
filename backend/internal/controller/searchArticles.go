package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetSearchResultArticles(c *gin.Context, articleAlias string) {
	article, err := h.articlesService.ReadArticle(articleAlias)
	if err != nil {
		c.JSON(http.StatusNotFound, &Error{
			Code:    int32(http.StatusNotFound),
			Type:    http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &Article{
		Id:            article.ID.String(),
		CreatedAt:     float32(article.CreatedAt.Unix()),
		UpdatedAt:     float32(article.UpdatedAt.Unix()),
		AuthorAddress: article.AuthorAddress,
		Title:         article.Title,
		Content:       article.Content,
		Alias:         &article.Alias,
	})
}
