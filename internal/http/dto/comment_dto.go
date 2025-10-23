package dto

import "belajar-go/internal/core/domain"

type CommentDTO struct {
	ID        uint   `json:"id"`
	PostID    uint   `json:"post_id"`
	UserID    uint   `json:"user_id"`
	ParentID  *uint  `json:"parent_id,omitempty"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewCommentDTO(d *domain.Comment) *CommentDTO {
	if d == nil {
		return nil
	}
	return &CommentDTO{
		ID:        d.ID,
		PostID:    d.PostID,
		UserID:    d.UserID,
		ParentID:  d.ParentID,
		Body:      d.Body,
		CreatedAt: d.CreatedAt.Format(timeLayout),
		UpdatedAt: d.UpdatedAt.Format(timeLayout),
	}
}

type CommentViewDTO struct {
	CommentDTO
	AuthorUsername string `json:"author_username"`
	LikeCount      int64  `json:"like_count"`
	LikedByMe      bool   `json:"liked_by_me"`
	ReplyCount     int64  `json:"reply_count"`
}

func NewCommentViewDTO(v *domain.CommentView) *CommentViewDTO {
	if v == nil {
		return nil
	}
	return &CommentViewDTO{
		CommentDTO: CommentDTO{
			ID:        v.ID,
			PostID:    v.PostID,
			UserID:    v.UserID,
			ParentID:  v.ParentID,
			Body:      v.Body,
			CreatedAt: v.CreatedAt.Format(timeLayout),
			UpdatedAt: v.UpdatedAt.Format(timeLayout),
		},
		AuthorUsername: v.AuthorUsername,
		LikeCount:      v.LikeCount,
		LikedByMe:      v.LikedByMe,
		ReplyCount:     v.ReplyCount,
	}
}

func NewCommentViewDTOs(vs []domain.CommentView) []CommentViewDTO {
	out := make([]CommentViewDTO, 0, len(vs))
	for i := range vs {
		out = append(out, *NewCommentViewDTO(&vs[i]))
	}
	return out
}
