package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ThreadLike struct {
	UserID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	ThreadID uuid.UUID `gorm:"type:uuid;primaryKey"`

	User   User   `gorm:"foreignKey:UserID;references:ID"`
	Thread Thread `gorm:"foreignKey:ThreadID;references:ID"`
}

func GetThreadLikeByThreadID(db *gorm.DB, threadID uuid.UUID) ([]ThreadLike, error) {
	var threadLikes []ThreadLike

	if err := db.Where("thread_id = ?", threadID).Find(&threadLikes).Error; err != nil {
		return threadLikes, err
	}

	return threadLikes, nil
}

func GetThreadLikeByUserIDAndThreadID(db *gorm.DB, userID, threadID uuid.UUID) (ThreadLike, error) {
	var threadLike ThreadLike

	if err := db.Where("user_id = ? AND thread_id = ?", userID, threadID).First(&threadLike).Error; err != nil {
		return threadLike, err
	}

	return threadLike, nil
}

func CreateThreadLike(db *gorm.DB, userID, threadID uuid.UUID) error {
	threadLike := ThreadLike{
		UserID:   userID,
		ThreadID: threadID,
	}

	if err := db.Create(&threadLike).Error; err != nil {
		return err
	}

	return nil
}

func DeleteThreadLike(db *gorm.DB, userID, threadID uuid.UUID) error {
	if err := db.Where("user_id = ? AND thread_id = ?", userID, threadID).Delete(&ThreadLike{}).Error; err != nil {
		return err
	}

	return nil
}
