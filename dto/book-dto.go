package dto

type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required,min=1"`
	Description string `json:"description" form:"description" binding:"required,min=1"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required,min=1"`
	Description string `json:"description" form:"description" binding:"required,min=1"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}