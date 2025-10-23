package gormrepo

import (
	"context"
	"errors"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"

	"gorm.io/gorm"
)

type retweetRepo struct{ db *gorm.DB }

func NewRetweetRepository(db *gorm.DB) ports.RetweetRepository { return &retweetRepo{db: db} }

func (r *retweetRepo) Create(ctx context.Context, d *domain.Retweet) error {
	m := toModelRetweet(d)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	d.ID = m.ID
	return nil
}

func (r *retweetRepo) Delete(ctx context.Context, userID, originalPostID uint) error {
	return r.db.WithContext(ctx).Where("user_id=? AND original_post_id=?", userID, originalPostID).
		Delete(&Retweet{}).Error
}

func (r *retweetRepo) FindByUserAndPost(ctx context.Context, userID, originalPostID uint) (*domain.Retweet, error) {
	var m Retweet
	if err := r.db.WithContext(ctx).
		Where("user_id=? AND original_post_id=?", userID, originalPostID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainRetweet(&m), nil
}

func (r *retweetRepo) ListByUser(ctx context.Context, userID uint) ([]domain.Retweet, error) {
	var ms []Retweet
	if err := r.db.WithContext(ctx).Where("user_id=?", userID).
		Order("id DESC").Find(&ms).Error; err != nil {
		return nil, err
	}
	out := make([]domain.Retweet, 0, len(ms))
	for i := range ms {
		out = append(out, *toDomainRetweet(&ms[i]))
	}
	return out, nil
}

func (r *retweetRepo) CountByPost(ctx context.Context, postID uint) (int64, error) {
	var cnt int64
	err := r.db.WithContext(ctx).Model(&Retweet{}).
		Where("original_post_id=?", postID).Count(&cnt).Error
	return cnt, err
}
