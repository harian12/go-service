package dto

// BookUpdateDTO is a model that used by client when PUT from /book/update url
type BookUpdateDTO struct {
	ID          uint64 `json:"id" from:"id" binding:"required"`
	Title       string `json:"title" from:"title" binding:"required"`
	Description string `json:"descr" from:"desc" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" from:"user_id,omitempty"`
}

// BookCreateDTO is a model that used by client when POST from /book/create url
type BookCreateDTO struct {
	ID          uint64 `json:"id" from:"id" binding:"required"`
	Title       string `json:"title" from:"title" binding:"required"`
	Description string `json:"descr" from:"desc" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" from:"user_id,omitempty"`
}
