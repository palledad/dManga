package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/palledad/dManga/backend/internal/models"
)

func (h *Handler) PostArticle(c *gin.Context) {
	articleReq := &PostArticleEntry{}
	err := c.BindJSON(&articleReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}
	tagUUIDs := make([]uuid.UUID, len(articleReq.Tags))
	for i, tagStr := range articleReq.Tags {
		tagUUID, err := uuid.Parse(tagStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, &Error{
				Code:    int32(http.StatusBadRequest),
				Type:    http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return
		}
		tagUUIDs[i] = tagUUID
	}
	article, err := h.articlesService.CreateArticle(&models.Article{
		ID:            uuid.New(),
		Title:         articleReq.Title,
		Content:       articleReq.Content,
		Alias:         *articleReq.Alias,
		AuthorAddress: articleReq.AuthorAddress,
	}, tagUUIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
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

func (h *Handler) GetArticle(c *gin.Context, articleAlias string) {
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

func (h *Handler) UpdateArticle(c *gin.Context, articleId string) {
	articleReq := &PostArticleEntry{}
	err := c.BindJSON(&articleReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}
	articleUUID, err := uuid.Parse(articleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}
	tagUUIDs := make([]uuid.UUID, len(articleReq.Tags))
	for i, tagStr := range articleReq.Tags {
		tagUUID, err := uuid.Parse(tagStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, &Error{
				Code:    int32(http.StatusBadRequest),
				Type:    http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return
		}
		tagUUIDs[i] = tagUUID
	}

	article, err := h.articlesService.UpdateArticle(articleUUID, articleReq.Title, articleReq.Content, tagUUIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Error{
			Code:    int32(http.StatusInternalServerError),
			Type:    http.StatusText(http.StatusInternalServerError),
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

func (h *Handler) DeleteArticle(c *gin.Context, articleId string) {
	uuid, err := uuid.Parse(articleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}
	article, err := h.articlesService.DeleteArticle(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Error{
			Code:    int32(http.StatusInternalServerError),
			Type:    http.StatusText(http.StatusInternalServerError),
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
