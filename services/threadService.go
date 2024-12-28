package services

import (
	"bone-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ThreadService struct {
	db *gorm.DB
}

func NewThreadService(db *gorm.DB) *ThreadService {
	return &ThreadService{db: db}
}

type ThreadResult struct {
	Thread  models.Thread
	IsLiked bool
	ThreadLikeCount int
}

func GetAllThread(db *gorm.DB, userID uuid.UUID) ([]ThreadResult, error) {
	threads, err := models.GetAllThread(db)

	if err != nil {
		return nil, err
	}

	var result []ThreadResult

	for _, thread := range threads {
		var isLiked bool

		_, err := models.GetThreadLikeByUserIDAndThreadID(db, userID, thread.ID)

		if err != nil {
			isLiked = false
		} else {
			isLiked = true
		}

		threadLikes, err := models.GetThreadLikeByThreadID(db, thread.ID)

		if err != nil {
			return nil, err
		}

		threadLikeCount := len(threadLikes)

		result = append(result, ThreadResult{Thread: thread, IsLiked: isLiked, ThreadLikeCount: threadLikeCount})
	}

	return result, nil
}
