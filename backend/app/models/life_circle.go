package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LifeCircle struct {
	gorm.Model
	HealthAndBody           int
	LoveAndRelationships    int
	FamilyAndFriends        int
	PersonalGrowth          int
	CareerAndFinance        int
	JoyAndRelax             int
	PhysicalEnvironment     int
	EmotionsAndFullfillment int

	UserID uuid.UUID `gorm:"type:uuid;not null;unique"`
}
