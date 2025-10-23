package ports

import (
	"belajar-go/internal/core/domain"
	"context"
)

type RetweetRepository interface {
	Create(ctx context.Context, r *domain.Retweet) error
	Delete(ctx context.Context, userID, originalPostID uint) error
	FindByUserAndPost(ctx context.Context, userID, originalPostID uint) (*domain.Retweet, error)

	ListByUser(ctx context.Context, userID uint) ([]domain.Retweet, error)
	CountByPost(ctx context.Context, postID uint) (int64, error)
}
