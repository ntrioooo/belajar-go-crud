package ports

import (
	"belajar-go/internal/core/domain"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, u *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindByID(ctx context.Context, id uint) (*domain.User, error)

	UpdateByID(ctx context.Context, id uint, fields map[string]any) error

	// untuk validasi unik saat update (exclude diri sendiri)
	ExistsEmailExcept(ctx context.Context, email string, exceptID uint) (bool, error)
	ExistsUsernameExcept(ctx context.Context, username string, exceptID uint) (bool, error)
}
