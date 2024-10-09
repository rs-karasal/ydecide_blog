package dto

// The request DTO for both register and login
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
