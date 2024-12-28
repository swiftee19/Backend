package services

import (
	"bone-backend/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ThreadLikeService struct {
	db *gorm.DB
}

func NewThreadLikeService(db *gorm.DB) *ThreadLikeService {
	return &ThreadLikeService{db: db}
}

func GetThreadLikeByUserIDAndThreadID(db *gorm.DB, userID, threadID uuid.UUID) (models.ThreadLike, error) {
	result, err := models.GetThreadLikeByUserIDAndThreadID(db, userID, threadID)

	if err != nil {
		return models.ThreadLike{}, err
	}

	return result, nil
}

func ThreadLike(db *gorm.DB, userID, threadID uuid.UUID) error {
	_, err := GetThreadLikeByUserIDAndThreadID(db, userID, threadID)

	fmt.Println(err)

	if err == nil {
		models.DeleteThreadLike(db, userID, threadID)
	} else {
		models.CreateThreadLike(db, userID, threadID)
	}

	return nil
}
