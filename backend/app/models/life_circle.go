package models

import (
	"github.com/google/uuid"
)

type LifeCircle struct {
	UUIDModel
	HealthAndBody           int `gorm:"check:health_and_body >= 0 AND health_and_body <= 10" json:"health_and_body"`
	LoveAndRelationships    int `gorm:"check:love_and_relationships >= 0 AND love_and_relationships <= 10" json:"love_and_relationships"`
	FamilyAndFriends        int `gorm:"check:family_and_friends >= 0 AND family_and_friends <= 10" json:"family_and_friends"`
	PersonalGrowth          int `gorm:"check:personal_growth >= 0 AND personal_growth <= 10" json:"personal_growth"`
	CareerAndFinance        int `gorm:"check:career_and_finance >= 0 AND career_and_finance <= 10" json:"career_and_finance"`
	JoyAndRelax             int `gorm:"check:joy_and_relax >= 0 AND joy_and_relax <= 10" json:"joy_and_relax"`
	PhysicalEnvironment     int `gorm:"check:physical_environment >= 0 AND physical_environment <= 10" json:"physical_environment"`
	EmotionsAndFullfillment int `gorm:"check:emotions_and_fullfillment >= 0 AND emotions_and_fullfillment <= 10" json:"emotions_and_fullfillment"`

	UserID uuid.UUID `gorm:"type:uuid;not null;unique" json:"user_id"`
}
