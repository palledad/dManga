package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID            uuid.UUID
	Title         string
	Content       string
	AuthorAddress string
	Alias         string
}

type ArticleModel struct {
}

func NewArticleModel() *ArticleModel {
	return &ArticleModel{}
}

func (m *ArticleModel) CreateArticle(db *gorm.DB, article *Article) (*Article, error) {
	if result := db.Create(article); result.Error != nil {
		return nil, result.Error
	}

	return article, nil
}

func (m *ArticleModel) ReadArticle(db *gorm.DB, alias string) (*Article, error) {
	target := &Article{}
	if result := db.Where("alias = ?", alias).First(&target); result.Error != nil {
		return nil, result.Error
	}
	return target, nil
}

func (m *ArticleModel) UpdateArticle(db *gorm.DB, articleId uuid.UUID, title string, content string) (*Article, error) {
	target := &Article{ID: articleId}
	if result := db.Model(target).Updates(Article{Title: title, Content: content}); result.Error != nil {
		return nil, result.Error
	}
	db.First(&target)
	return target, nil
}

func (m *ArticleModel) DeleteArticle(db *gorm.DB, articleId uuid.UUID) (*Article, error) {
	target := &Article{ID: articleId}
	if result := db.Delete(&target); result.Error != nil {
		return nil, result.Error
	}
	db.Unscoped().Find(&target)
	return target, nil
}
