package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
// }
