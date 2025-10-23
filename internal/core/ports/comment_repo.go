package ports

import (
	"belajar-go/internal/core/domain"
	"context"
)

type CommentRepository interface {
	Create(ctx context.Context, c *domain.Comment) error
	FindByID(ctx context.Context, id uint) (*domain.Comment, error)
	ListByPost(ctx context.Context, postID uint) ([]domain.Comment, error)
	ListReplies(ctx context.Context, parentID uint) ([]domain.Comment, error)
	Update(ctx context.Context, c *domain.Comment) error
	Delete(ctx context.Context, id uint) error

	CountByPost(ctx context.Context, postID uint) (int64, error)
	CountReplies(ctx context.Context, parentID uint) (int64, error)

	Like(ctx context.Context, userID, commentID uint) error
	Unlike(ctx context.Context, userID, commentID uint) error
	IsLiked(ctx context.Context, userID, commentID uint) (bool, error)
	CountLikes(ctx context.Context, commentID uint) (int64, error)
}
