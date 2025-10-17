package gormrepo

import (
	"context"
	"errors"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"

	"gorm.io/gorm"
)

type postRepo struct{ db *gorm.DB }

func NewPostRepository(db *gorm.DB) ports.PostRepository { return &postRepo{db: db} }

func (r *postRepo) Create(ctx context.Context, p *domain.Post) error {
	m := toModelPost(p)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	p.ID = m.ID
	return nil
}

func (r *postRepo) FindByID(ctx context.Context, id uint) (*domain.Post, error) {
	var m Post
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainPost(&m), nil
}

func (r *postRepo) List(ctx context.Context) ([]domain.Post, error) {
	var ms []Post
	if err := r.db.WithContext(ctx).Order("id DESC").Find(&ms).Error; err != nil {
		return nil, err
	}
	out := make([]domain.Post, 0, len(ms))
	for i := range ms {
		out = append(out, *toDomainPost(&ms[i]))
	}
	return out, nil
}

func (r *postRepo) Update(ctx context.Context, p *domain.Post) error {
	m := toModelPost(p)
	return r.db.WithContext(ctx).Model(&Post{}).Where("id=?", m.ID).Updates(map[string]any{"title": m.Title, "body": m.Body}).Error
}

func (r *postRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Post{}, id).Error
}
