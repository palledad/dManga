package services

import (
	"github.com/google/uuid"
	"github.com/palledad/dManga/backend/internal/models"
	"gorm.io/gorm"
)

type ArticleService struct {
	articleModel      *models.ArticleModel
	articlesTagsModel *models.ArticlesTagsModel
	db                *gorm.DB
}

func NewArticlesService(db *gorm.DB, articleModel *models.ArticleModel) *ArticleService {
	return &ArticleService{
		articleModel: articleModel,
		db:           db,
	}
}

func (a *ArticleService) CreateArticle(article *models.Article, tagIDs []uuid.UUID) (*models.Article, error) {
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if _, err := a.articleModel.CreateArticle(tx, article); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := a.articlesTagsModel.RecordTags(tx, article.ID, tagIDs); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
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

func (a *ArticleService) UpdateArticle(articleId uuid.UUID, title string, content string, tagIDs []uuid.UUID) (*models.Article, error) {
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	article, err := a.articleModel.UpdateArticle(a.db, articleId, title, content)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := a.articlesTagsModel.UpdateTags(tx, articleId, tagIDs); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleService) DeleteArticle(articleId uuid.UUID) (*models.Article, error) {
	return a.articleModel.DeleteArticle(a.db, articleId)
}
