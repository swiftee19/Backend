package models

import "github.com/google/uuid"

type ThreadTag struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `json:"name"`
}
