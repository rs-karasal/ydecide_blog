package dto

type UpdateLifeCircleRequest struct {
	HealthAndBody           int `json:"health_and_body" validate:"min=0,max=10"`
	LoveAndRelationships    int `json:"love_and_relationships" validate:"min=0,max=10"`
	FamilyAndFriends        int `json:"family_and_friends" validate:"min=0,max=10"`
	PersonalGrowth          int `json:"personal_growth" validate:"min=0,max=10"`
	CareerAndFinance        int `json:"career_and_finance" validate:"min=0,max=10"`
	JoyAndRelax             int `json:"joy_and_relax" validate:"min=0,max=10"`
	PhysicalEnvironment     int `json:"physical_environment" validate:"min=0,max=10"`
	EmotionsAndFullfillment int `json:"emotions_and_fullfillment" validate:"min=0,max=10"`
}
