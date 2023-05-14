package services

import (
	"github.com/google/uuid"
	"github.com/palledad/dManga/backend/internal/models"
	"gorm.io/gorm"
)

type TagService struct {
	tagModel *models.TagModel
	db       *gorm.DB
}

func NewTagService(db *gorm.DB, tagModel *models.TagModel) *TagService {
	return &TagService{
		tagModel: tagModel,
		db:       db,
	}
}

func (s *TagService) CreateTag(tag *models.Tag) error {
	return s.tagModel.CreateTag(s.db, tag)
}

func (s *TagService) GetTagByID(tagID uuid.UUID) (*models.Tag, error) {
	return s.tagModel.GetTagByID(s.db, tagID)
}

func (s *TagService) UpdateTag(tagID uuid.UUID, tagName string) (*models.Tag, error) {
	return s.tagModel.UpdateTag(s.db, tagID, tagName)
}

func (s *TagService) DeleteTag(tagID uuid.UUID) (*models.Tag, error) {
	return s.tagModel.DeleteTag(s.db, tagID)
}
