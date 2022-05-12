package dto

// UserCreateDTO is a model that used by client when POST from /user/create url
type UserCreateDTO struct {
	ID       uint64 `json:"id" from:"id" binding:"required"`
	Name     string `json:"name" from:"name" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required"`
	Password string `json:"password,omitempty" from:"password,omitempty" binding:"required" validate:"min:6"`
}

// UserUpdateDTO is used when update a user
type UserUpdateDTO struct {
	ID       uint64 `json:"id" from:"id" binding:"required"`
	Name     string `json:"name" from:"name" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" from:"password,omitempty" validate:"min:6"`
}
