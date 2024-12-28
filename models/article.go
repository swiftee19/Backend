package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	PostedAt    time.Time `gorm:"type:timestamp;default:now();autoUpdateTime"`
}

func GetAllArticle(db *gorm.DB) ([]Article, error) {
	var articles []Article

	result := db.Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}

func GetLatestArticle(db *gorm.DB) (Article, error) {
	var article Article

	result := db.Order("posted_at desc").First(&article)

	if result.Error != nil {
		return Article{}, result.Error
	}

	return article, nil
}
