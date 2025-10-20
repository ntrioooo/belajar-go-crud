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

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var m User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainUser(&m), nil
}

func (r *userRepo) UpdateByID(ctx context.Context, id uint, fields map[string]any) error {
	return r.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Updates(fields).Error
}

func (r *userRepo) ExistsEmailExcept(ctx context.Context, email string, exceptID uint) (bool, error) {
	var cnt int64
	if err := r.db.WithContext(ctx).Model(&User{}).
		Where("email = ? AND id <> ?", email, exceptID).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (r *userRepo) ExistsUsernameExcept(ctx context.Context, username string, exceptID uint) (bool, error) {
	var cnt int64
	if err := r.db.WithContext(ctx).Model(&User{}).
		Where("username = ? AND id <> ?", username, exceptID).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}
