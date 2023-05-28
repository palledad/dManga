package services

import (
	"github.com/palledad/dManga/backend/internal/models"
	"gorm.io/gorm"
)

type SearchArticlesService struct {
	articleModel *models.ArticleModel
	db           *gorm.DB
}

func NewSearchArticlesService(db *gorm.DB, articleModel *models.ArticleModel) *ArticleService {
	return &ArticleService{
		articleModel: articleModel,
		db:           db,
	}
}
func (a *ArticleService) GetSearchResultArticles(alias string) (*models.Article, error) {
	article, err := a.articleModel.ReadArticle(a.db, alias)
	if err != nil {
		return nil, err
	}
	return article, nil
}
