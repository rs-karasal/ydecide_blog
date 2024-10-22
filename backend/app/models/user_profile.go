package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserProfile struct {
	UUIDModel
	Name        sql.NullString `json:"name"`
	BirthDate   sql.NullString `json:"birth_date"`
	Photo       sql.NullString `json:"photo"`
	Description sql.NullString `json:"description"`
	Profession  sql.NullString `json:"profession"`
	Languages   sql.NullString `json:"languages"`
	Location    sql.NullString `json:"location"`

	UserID uuid.UUID `gorm:"type:uuid;not null;unique" json:"user_id"`
}
