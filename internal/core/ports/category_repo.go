package ports

import (
	"belajar-go/internal/core/domain"
	"context"
)

type CategoryRepository interface {
	Create(ctx context.Context, p *domain.Category) error
	FindByID(ctx context.Context, id uint) (*domain.Category, error)
	List(ctx context.Context) ([]domain.Category, error)
	Update(ctx context.Context, p *domain.Category) error
	Delete(ctx context.Context, id uint) error
}
