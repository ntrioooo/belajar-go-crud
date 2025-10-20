package services

import (
	"context"
	"errors"
	"strings"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"
)

type categoryService struct {
	categories ports.CategoryRepository
}

func NewCategoryService(categories ports.CategoryRepository) ports.CategoryService {
	return &categoryService{categories: categories}
}

func (s *categoryService) Create(ctx context.Context, name string) (*domain.Category, error) {
	name = strings.TrimSpace(name)
	if len(name) < 3 {
		return nil, errors.New("name too short")
	}
	c := &domain.Category{Name: name}
	if err := s.categories.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

// func (s *postService) GetByID(ctx context.Context, id uint) (*domain.Post, error) {
// 	return s.posts.FindByID(ctx, id)
// }

// func (s *postService) List(ctx context.Context) ([]domain.Post, error) {
// 	return s.posts.List(ctx)
// }

func (s *categoryService) Update(ctx context.Context, id uint, name string) (*domain.Category, error) {
	c, err := s.categories.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	c.Name = name
	if err := s.categories.Update(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *categoryService) Delete(ctx context.Context, id uint) error {
	return s.categories.Delete(ctx, id)
}

func (s *categoryService) List(ctx context.Context) ([]domain.CategoryView, error) {
	cats, err := s.categories.List(ctx)
	if err != nil {
		return nil, err
	}

	views := make([]domain.CategoryView, 0, len(cats))
	for i := range cats {
		views = append(views, domain.CategoryView{Category: cats[i]})
	}

	return views, nil
}

func (s *categoryService) GetByID(ctx context.Context, id uint) (*domain.CategoryView, error) {
	p, err := s.categories.FindByID(ctx, id)
	if err != nil || p == nil {
		return nil, err
	}
	view := &domain.CategoryView{
		Category: *p,
	}
	return view, nil
}
