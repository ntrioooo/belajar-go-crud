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

// Like/Unlike
func (r *postRepo) Like(ctx context.Context, userID, postID uint) error {
	// gunakan INSERT IGNORE-like behavior: cek dulu
	var cnt int64
	if err := r.db.WithContext(ctx).Model(&Like{}).
		Where("user_id=? AND post_id=?", userID, postID).
		Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return nil
	}
	return r.db.WithContext(ctx).Create(&Like{UserID: userID, PostID: postID}).Error
}
func (r *postRepo) Unlike(ctx context.Context, userID, postID uint) error {
	return r.db.WithContext(ctx).Where("user_id=? AND post_id=?", userID, postID).Delete(&Like{}).Error
}
func (r *postRepo) IsLiked(ctx context.Context, userID, postID uint) (bool, error) {
	var cnt int64
	if err := r.db.WithContext(ctx).Model(&Like{}).
		Where("user_id=? AND post_id=?", userID, postID).
		Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}
func (r *postRepo) CountLikes(ctx context.Context, postID uint) (int64, error) {
	var cnt int64
	err := r.db.WithContext(ctx).Model(&Like{}).Where("post_id=?", postID).Count(&cnt).Error
	return cnt, err
}
func (r *postRepo) BatchAuthorUsernames(ctx context.Context, userIDs []uint) (map[uint]string, error) {
	if len(userIDs) == 0 {
		return map[uint]string{}, nil
	}
	var users []User
	if err := r.db.WithContext(ctx).Model(&User{}).
		Where("id IN ?", userIDs).Select("id, username").Find(&users).Error; err != nil {
		return nil, err
	}
	m := make(map[uint]string, len(users))
	for i := range users {
		m[users[i].ID] = users[i].Username
	}
	return m, nil
}
