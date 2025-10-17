package ports

import (
	"belajar-go/internal/core/domain"
	"context"
)

type PostRepository interface {
	Create(ctx context.Context, p *domain.Post) error
	FindByID(ctx context.Context, id uint) (*domain.Post, error)
	List(ctx context.Context) ([]domain.Post, error)
	Update(ctx context.Context, p *domain.Post) error
	Delete(ctx context.Context, id uint) error
}
