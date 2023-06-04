package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleTag struct {
	gorm.Model
	ArticleID uuid.UUID
	TagID     uuid.UUID
}

type ArticlesTagsModel struct {
}

func NewArticlesTags() *ArticlesTagsModel {
	return &ArticlesTagsModel{}
}

func (m *ArticlesTagsModel) RecordTags(db *gorm.DB, articleID uuid.UUID, tagIDs []uuid.UUID) error {
	recordingTags := make([]ArticleTag, len(tagIDs))
	for i, tagID := range tagIDs {
		articleTag := ArticleTag{
			ArticleID: articleID,
			TagID:     tagID,
		}
		recordingTags[i] = articleTag
	}
	return db.Create(recordingTags).Error
}

func (m *ArticlesTagsModel) UpdateTags(db *gorm.DB, articleID uuid.UUID, tagIDs []uuid.UUID) error {
	if err := db.Unscoped().Delete(&ArticleTag{}, "artilce_id = ?", articleID).Error; err != nil {
		return err
	}
	recordingTags := make([]ArticleTag, len(tagIDs))
	for i, tagID := range tagIDs {
		articleTag := ArticleTag{
			ArticleID: articleID,
			TagID:     tagID,
		}
		recordingTags[i] = articleTag
	}
	return db.Create(recordingTags).Error
}
