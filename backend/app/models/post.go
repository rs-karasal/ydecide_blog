package models

import (
	"github.com/google/uuid"
)

type Post struct {
	UUIDModel
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`

	UserID *uuid.UUID `gorm:"null" json:"author_id"`
}
