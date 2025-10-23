package gormrepo

import (
	"context"
	"errors"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"

	"gorm.io/gorm"
)

type commentRepo struct{ db *gorm.DB }

func NewCommentRepository(db *gorm.DB) ports.CommentRepository { return &commentRepo{db: db} }

// ===== CRUD =====

func (r *commentRepo) Create(ctx context.Context, c *domain.Comment) error {
	m := toModelComment(c)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	c.ID = m.ID
	return nil
}

func (r *commentRepo) FindByID(ctx context.Context, id uint) (*domain.Comment, error) {
	var m Comment
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainComment(&m), nil
}

func (r *commentRepo) ListByPost(ctx context.Context, postID uint) ([]domain.Comment, error) {
	var ms []Comment
	if err := r.db.WithContext(ctx).
		Where("post_id = ? AND parent_id IS NULL", postID).
		Order("id ASC").Find(&ms).Error; err != nil {
		return nil, err
	}
	out := make([]domain.Comment, 0, len(ms))
	for i := range ms {
		out = append(out, *toDomainComment(&ms[i]))
	}
	return out, nil
}

func (r *commentRepo) ListReplies(ctx context.Context, parentID uint) ([]domain.Comment, error) {
	var ms []Comment
	if err := r.db.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Order("id ASC").Find(&ms).Error; err != nil {
		return nil, err
	}
	out := make([]domain.Comment, 0, len(ms))
	for i := range ms {
		out = append(out, *toDomainComment(&ms[i]))
	}
	return out, nil
}

func (r *commentRepo) Update(ctx context.Context, c *domain.Comment) error {
	m := toModelComment(c)
	return r.db.WithContext(ctx).Model(&Comment{}).
		Where("id = ?", m.ID).
		Updates(map[string]any{"body": m.Body}).Error
}

func (r *commentRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Comment{}, id).Error
}

// ===== Counts =====

func (r *commentRepo) CountByPost(ctx context.Context, postID uint) (int64, error) {
	var cnt int64
	err := r.db.WithContext(ctx).Model(&Comment{}).
		Where("post_id = ?", postID).Count(&cnt).Error
	return cnt, err
}
func (r *commentRepo) CountReplies(ctx context.Context, parentID uint) (int64, error) {
	var cnt int64
	err := r.db.WithContext(ctx).Model(&Comment{}).
		Where("parent_id = ?", parentID).Count(&cnt).Error
	return cnt, err
}

// ===== Likes for comments =====

func (r *commentRepo) Like(ctx context.Context, userID, commentID uint) error {
	var cnt int64
	if err := r.db.WithContext(ctx).Model(&CommentLike{}).
		Where("user_id=? AND comment_id=?", userID, commentID).
		Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return nil
	}
	return r.db.WithContext(ctx).Create(&CommentLike{UserID: userID, CommentID: commentID}).Error
}

func (r *commentRepo) Unlike(ctx context.Context, userID, commentID uint) error {
	return r.db.WithContext(ctx).Where("user_id=? AND comment_id=?", userID, commentID).
		Delete(&CommentLike{}).Error
}

func (r *commentRepo) IsLiked(ctx context.Context, userID, commentID uint) (bool, error) {
	var cnt int64
	if err := r.db.WithContext(ctx).Model(&CommentLike{}).
		Where("user_id=? AND comment_id=?", userID, commentID).
		Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (r *commentRepo) CountLikes(ctx context.Context, commentID uint) (int64, error) {
	var cnt int64
	err := r.db.WithContext(ctx).Model(&CommentLike{}).
		Where("comment_id=?", commentID).Count(&cnt).Error
	return cnt, err
}
