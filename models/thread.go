package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Thread struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UploaderID uuid.UUID `gorm:"type:uuid;not null;index"`
	Content    string    `json:"content"`

	// Foreign key relation to User
	Uploader User `gorm:"foreignKey:UploaderID;references:ID"`
}

func GetAllThread(db *gorm.DB) ([]Thread, error) {
	var threads []Thread

	result := db.Preload("Uploader").Find(&threads)

	if result.Error != nil {
		return nil, result.Error
	}

	return threads, nil
}
