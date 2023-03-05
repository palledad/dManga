package services

import (
	"github.com/google/uuid"
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
	if _, err := a.articleModel.CreateArticle(a.db, article); err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleService) ReadArticle(alias string) (*models.Article, error) {
	article, err := a.articleModel.ReadArticle(a.db, alias)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleService) UpdateArticle(articleId uuid.UUID, title string, content string) (*models.Article, error) {
	tx := a.db.Begin()
	articleResp, err := a.articleModel.UpdateArticle(tx, articleId, title, content)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return articleResp, nil
}
