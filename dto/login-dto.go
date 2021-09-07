package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required" validate:"min:6"`
}
