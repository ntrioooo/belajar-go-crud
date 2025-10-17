package ports

import (
	"belajar-go/internal/core/domain"
	"context"
)

type AuthService interface {
	Signup(ctx context.Context, email, password string) (*domain.User, error)
	Login(ctx context.Context, email, password string) (token string, user *domain.User, err error)
}

type PostService interface {
	Create(ctx context.Context, authorID uint, title, body string) (*domain.Post, error)
	GetByID(ctx context.Context, id uint) (*domain.Post, error)
	List(ctx context.Context) ([]domain.Post, error)
	Update(ctx context.Context, id uint, title, body string) (*domain.Post, error)
	Delete(ctx context.Context, id uint) error
}
