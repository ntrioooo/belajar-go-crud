package gormrepo

import (
	"context"
	"errors"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"

	"gorm.io/gorm"
)

type CategoryRepo struct{ db *gorm.DB }

func NewCategoryRepository(db *gorm.DB) ports.CategoryRepository { return &CategoryRepo{db: db} }

func (r *CategoryRepo) Create(ctx context.Context, c *domain.Category) error {
	m := toModelCategory(c)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	c.ID = m.ID
	return nil
}

func (r *CategoryRepo) FindByID(ctx context.Context, id uint) (*domain.Category, error) {
	var m Category
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainCategory(&m), nil
}

func (r *CategoryRepo) List(ctx context.Context) ([]domain.Category, error) {
	var ms []Category
	if err := r.db.WithContext(ctx).Order("id DESC").Find(&ms).Error; err != nil {
		return nil, err
	}
	out := make([]domain.Category, 0, len(ms))
	for i := range ms {
		out = append(out, *toDomainCategory(&ms[i]))
	}
	return out, nil
}

func (r *CategoryRepo) Update(ctx context.Context, c *domain.Category) error {
	m := toModelCategory(c)
	return r.db.WithContext(ctx).Model(&Category{}).Where("id=?", m.ID).Updates(map[string]any{"name": m.Name}).Error
}

func (r *CategoryRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Category{}, id).Error
}
