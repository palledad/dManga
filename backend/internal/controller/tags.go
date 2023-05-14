package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/palledad/dManga/backend/internal/models"
)

// (POST /tags)
func (h Handler) PostTag(c *gin.Context) {
	tagReq := &PostTagEntry{}
	err := c.BindJSON(&tagReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}
	tag := &models.Tag{
		ID:   uuid.New(),
		Name: tagReq.Name,
	}
	if err := h.tagService.CreateTag(tag); err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Tag{
		Id:        tag.ID.String(),
		Name:      tag.Name,
		CreatedAt: float32(tag.CreatedAt.Unix()),
		UpdatedAt: float32(tag.UpdatedAt.Unix()),
	})
}

// (DELETE /tags/{tag_id})
func (h Handler) DeleteTag(c *gin.Context, tagId string) {
	uuid, err := uuid.Parse(tagId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}
	tag, err := h.tagService.DeleteTag(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, &Error{
			Code:    int32(http.StatusNotFound),
			Type:    http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &Tag{
		Id:        tag.ID.String(),
		Name:      tag.Name,
		CreatedAt: float32(tag.CreatedAt.Unix()),
		UpdatedAt: float32(tag.UpdatedAt.Unix()),
	})
}

// (GET /tags/{tag_id})
func (h Handler) GetTag(c *gin.Context, tagId string) {
	uuid, err := uuid.Parse(tagId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}
	tag, err := h.tagService.GetTagByID(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, &Error{
			Code:    int32(http.StatusNotFound),
			Type:    http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &Tag{
		Id:        tag.ID.String(),
		Name:      tag.Name,
		CreatedAt: float32(tag.CreatedAt.Unix()),
		UpdatedAt: float32(tag.UpdatedAt.Unix()),
	})
}

// (PATCH /tags/{tag_id})
func (h Handler) UpdateTag(c *gin.Context, tagId string) {
	tagReq := &PostTagEntry{}
	err := c.BindJSON(&tagReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}
	uuid, err := uuid.Parse(tagId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Error{
			Code:    int32(http.StatusBadRequest),
			Type:    http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}
	tag, err := h.tagService.UpdateTag(uuid, tagReq.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Error{
			Code:    int32(http.StatusInternalServerError),
			Type:    http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &Tag{
		Id:        tag.ID.String(),
		Name:      tag.Name,
		CreatedAt: float32(tag.CreatedAt.Unix()),
		UpdatedAt: float32(tag.UpdatedAt.Unix()),
	})
}
