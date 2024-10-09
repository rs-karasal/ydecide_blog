package models

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"not null" json:"username"`
	PasswordHash string `gorm:"not null" json:"-"`
	Posts        []Post `gorm:"foreignKey:AuthorID" json:"posts"`
}
