package dto

import "belajar-go/internal/core/domain"

// Untuk list/show post (gabungan post + info agregat)
type PostViewDTO struct {
	ID                 uint    `json:"id"`
	Title              string  `json:"title"`
	Body               string  `json:"body"`
	UserID             uint    `json:"user_id"`
	CategoryID         uint    `json:"category_id"`
	AuthorUsername     string  `json:"author_username"`
	CategoryName       string  `json:"category_name"`
	LikeCount          int64   `json:"like_count"`
	LikedByMe          bool    `json:"liked_by_me"`
	CommentCount       int64   `json:"comment_count"`
	RetweetCount       int64   `json:"retweet_count"`
	RepostedByUsername *string `json:"reposted_by_username,omitempty"`
	QuoteBody          *string `json:"quote_body,omitempty"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
	// (opsional) event_at utk timeline (isi waktu retweet, kalau ada)
	EventAt *string `json:"event_at,omitempty"`
}

func NewPostViewDTO(v *domain.PostView) *PostViewDTO {
	if v == nil {
		return nil
	}
	dto := &PostViewDTO{
		ID:                 v.ID,
		Title:              v.Title,
		Body:               v.Body,
		UserID:             v.UserID,
		CategoryID:         v.CategoryID,
		AuthorUsername:     v.AuthorUsername,
		CategoryName:       v.CategoryName,
		LikeCount:          v.LikeCount,
		LikedByMe:          v.LikedByMe,
		CommentCount:       v.CommentCount,
		RetweetCount:       v.RetweetCount,
		RepostedByUsername: v.RepostedByUsername,
		QuoteBody:          v.QuoteBody,
		CreatedAt:          v.CreatedAt.Format(timeLayout),
		UpdatedAt:          v.UpdatedAt.Format(timeLayout),
	}
	// kalau kamu tambahkan EventAt di domain.PostView, map juga ke sini.
	return dto
}

func NewPostViewDTOs(vs []domain.PostView) []PostViewDTO {
	out := make([]PostViewDTO, 0, len(vs))
	for i := range vs {
		out = append(out, *NewPostViewDTO(&vs[i]))
	}
	return out
}
