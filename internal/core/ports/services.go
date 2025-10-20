package ports

import (
	"belajar-go/internal/core/domain"
	"context"
)

type AuthService interface {
	Signup(ctx context.Context, email, username, password string) (*domain.User, error)
	Login(ctx context.Context, email, password string) (token string, user *domain.User, err error)
}

type PostService interface {
	Create(ctx context.Context, authorID uint, title, body string, categoryID uint) (*domain.Post, error)
	GetByID(ctx context.Context, viewerID, id uint) (*domain.PostView, error)
	List(ctx context.Context, viewerID uint) ([]domain.PostView, error)
	Update(ctx context.Context, id uint, title, body string, categoryID uint) (*domain.Post, error)
	Delete(ctx context.Context, id uint) error
	ToggleLike(ctx context.Context, userID, postID uint) (liked bool, likeCount int64, err error)
}

type UserService interface {
	GetMe(ctx context.Context, userID uint) (*domain.User, error)
	UpdateMe(ctx context.Context, userID uint, email, username *string) (*domain.User, error)
}

type CategoryService interface {
	Create(ctx context.Context, name string) (*domain.Category, error)
	GetByID(ctx context.Context, id uint) (*domain.CategoryView, error)
	List(ctx context.Context) ([]domain.CategoryView, error)
	Update(ctx context.Context, id uint, name string) (*domain.Category, error)
	Delete(ctx context.Context, id uint) error
}
