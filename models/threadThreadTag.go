package models

import "github.com/google/uuid"

type ThreadThreadTag struct {
	ThreadID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	ThreadTagID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Thread Thread    `gorm:"foreignKey:ThreadID;references:ID"`
	Tag    ThreadTag `gorm:"foreignKey:ThreadTagID;references:ID"`
}
