package dto

// RegisterDTO is a model that used by client when POST from /register url
type RegisterDTO struct {
	Name     string `json:"name" from:"name" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required"`
	Password string `json:"password" from:"password" binding:"required" validate:"min:6"`
}
