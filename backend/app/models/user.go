package models

type User struct {
	UUIDModel
	Username     string `gorm:"not null;uniqueIndex" json:"username"`
	PasswordHash string `gorm:"not null" json:"-"`

	Posts       []Post       `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL" json:"posts"`
	LifeCircle  *LifeCircle  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserProfile *UserProfile `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
