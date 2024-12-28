package services

import (
	"bone-backend/models"

	"gorm.io/gorm"
)

type ArticleService struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{db: db}
}

func GetAllArticle(db *gorm.DB) ([]models.Article, error) {
	var articles []models.Article

	result := db.Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}

func GetLatestArticle(db *gorm.DB) (models.Article, error) {
	article, err := models.GetLatestArticle(db)

	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}
