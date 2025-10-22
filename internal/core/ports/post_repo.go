package ports

import (
	"belajar-go/internal/core/domain"
	"context"
)

type PostRepository interface {
	Create(ctx context.Context, p *domain.Post) error
	FindByID(ctx context.Context, id uint) (*domain.Post, error)
	// GetByID(ctx context.Context, id uint) (*domain.Post, error)
	List(ctx context.Context) ([]domain.Post, error)
	Update(ctx context.Context, p *domain.Post) error
	Delete(ctx context.Context, id uint) error

	Like(ctx context.Context, userID, postID uint) error
	Unlike(ctx context.Context, userID, postID uint) error
	IsLiked(ctx context.Context, userID, postID uint) (bool, error)
	CountLikes(ctx context.Context, postID uint) (int64, error)
	BatchAuthorUsernames(ctx context.Context, userIDs []uint) (map[uint]string, error)
	BatchCategoryNames(ctx context.Context, ids []uint) (map[uint]string, error) // <-- tambahkan ini
}
