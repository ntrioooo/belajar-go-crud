package gormrepo

import (
	"context"
	"errors"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"

	"gorm.io/gorm"
)

type userRepo struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) ports.UserRepository { return &userRepo{db: db} }

func (r *userRepo) Create(ctx context.Context, u *domain.User) error {
	m := toModelUser(u)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	u.ID = m.ID
	return nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var m User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainUser(&m), nil
}

func (r *userRepo) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	var m User
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainUser(&m), nil
}
