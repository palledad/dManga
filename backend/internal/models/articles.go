package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID            uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
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

func (a *ArticleModel) ReadArticle(db *gorm.DB, alias string) (*Article, error) {
	target := &Article{Alias: alias}
	if result := db.First(&target); result.Error != nil {
		return nil, result.Error
	}
	return target, nil
}

func (m *ArticleModel) CreateArticle(db *gorm.DB, article *Article) error {
	fmt.Println(article.ID)
	if result := db.Create(article); result.Error != nil {
		return result.Error
	}
	return nil
}
