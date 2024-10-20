package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string    `gorm:"not null" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`

	Posts      []Post     `gorm:"foreignKey:AuthorID" json:"posts"`
	LifeCircle LifeCircle `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
