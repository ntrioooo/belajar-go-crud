package services

import (
	"context"
	"errors"
	"strings"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"
)

type postService struct {
	posts ports.PostRepository
}

func NewPostService(posts ports.PostRepository) ports.PostService {
	return &postService{posts: posts}
}

func (s *postService) Create(ctx context.Context, authorID uint, title, body string) (*domain.Post, error) {
	title = strings.TrimSpace(title)
	body = strings.TrimSpace(body)
	if authorID == 0 {
		return nil, errors.New("unauthorized")
	}
	if len(title) < 3 || len(body) < 3 {
		return nil, errors.New("title/body too short")
	}

	p := &domain.Post{Title: title, Body: body, UserID: authorID}
	if err := s.posts.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *postService) GetByID(ctx context.Context, id uint) (*domain.Post, error) {
	return s.posts.FindByID(ctx, id)
}

func (s *postService) List(ctx context.Context) ([]domain.Post, error) {
	return s.posts.List(ctx)
}

func (s *postService) Update(ctx context.Context, id uint, title, body string) (*domain.Post, error) {
	p, err := s.posts.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	p.Title, p.Body = title, body
	if err := s.posts.Update(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *postService) Delete(ctx context.Context, id uint) error {
	return s.posts.Delete(ctx, id)
}
