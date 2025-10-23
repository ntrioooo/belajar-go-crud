package dto

import (
	"belajar-go/internal/core/domain"
)

type UserDTO struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserDTO(d *domain.User) *UserDTO {
	if d == nil {
		return nil
	}
	return &UserDTO{
		ID:        d.ID,
		Email:     d.Email,
		Username:  d.Username,
		Role:      d.Role,
		CreatedAt: d.CreatedAt.Format(timeLayout),
		UpdatedAt: d.UpdatedAt.Format(timeLayout),
	}
}
