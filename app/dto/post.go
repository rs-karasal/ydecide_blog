package dto

type PostRequest struct {
	Title    string `json:"title" validate:"required"`
	AuthorID uint   `json:"author_id"`
}
