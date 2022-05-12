package dto

//LoginDTO is a model that used by client when POST from /login url
type LoginDTO struct {
	Email    string `json:"email" from:"email" binding:"required"`
	Password string `json:"password" from:"password" binding:"required" validate:"min:6"`
}
