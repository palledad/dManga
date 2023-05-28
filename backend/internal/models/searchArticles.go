package models

import (
	"gorm.io/gorm"
)

type SearchArticleModel struct {
}

func (m *ArticleModel) GetSearchResultArticles(db *gorm.DB) (*Article, error) {
	target := &Article{}
	if result := db.First(&target); result.Error != nil {
		return nil, result.Error
	}
	return target, nil
}
