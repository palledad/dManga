package services

import (
	"github.com/palledad/dManga/backend/internal/models"
	"gorm.io/gorm"
)

type ArticleService struct {
	articleModel *models.ArticleModel
	db           *gorm.DB
}

func NewArticlesService(db *gorm.DB, articleModel *models.ArticleModel) *ArticleService {
	return &ArticleService{
		articleModel: articleModel,
		db:           db,
	}
}

func (a *ArticleService) CreateArticle(article *models.Article) (*models.Article, error) {
	if err := a.articleModel.CreateArticle(a.db, article); err != nil {
		return nil, err
	}
	return article, nil
}
