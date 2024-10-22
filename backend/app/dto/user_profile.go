package dto

type UserProfileRequest struct {
	Name        string `json:"name"`
	BirthDate   string `json:"birth_date"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
	Profession  string `json:"profession"`
	Languages   string `json:"languages"`
	Location    string `json:"location"`
}
