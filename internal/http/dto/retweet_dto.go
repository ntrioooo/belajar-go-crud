package dto

import "belajar-go/internal/core/domain"

type PostSummaryDTO struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Body           string `json:"body"`
	AuthorUsername string `json:"author_username"`
	CategoryName   string `json:"category_name"`
}

type RetweetDTO struct {
	ID             uint    `json:"id"`
	UserID         uint    `json:"user_id"`
	OriginalPostID uint    `json:"original_post_id"`
	QuoteBody      *string `json:"quote_body,omitempty"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type RetweetViewDTO struct {
	RetweetDTO
	Original PostSummaryDTO `json:"original"`
}

func NewRetweetDTO(d *domain.Retweet) *RetweetDTO {
	if d == nil {
		return nil
	}
	return &RetweetDTO{
		ID:             d.ID,
		UserID:         d.UserID,
		OriginalPostID: d.OriginalPostID,
		QuoteBody:      d.QuoteBody,
		CreatedAt:      d.CreatedAt.Format(timeLayout),
		UpdatedAt:      d.UpdatedAt.Format(timeLayout),
	}
}

func NewRetweetViewDTO(v *domain.RetweetView) *RetweetViewDTO {
	if v == nil {
		return nil
	}
	return &RetweetViewDTO{
		RetweetDTO: *NewRetweetDTO(&v.Retweet),
		Original: PostSummaryDTO{
			ID:             v.Original.ID,
			Title:          v.Original.Title,
			Body:           v.Original.Body,
			AuthorUsername: v.Original.AuthorUsername,
			CategoryName:   v.Original.CategoryName,
		},
	}
}

func NewRetweetViewDTOs(vs []domain.RetweetView) []RetweetViewDTO {
	out := make([]RetweetViewDTO, 0, len(vs))
	for i := range vs {
		out = append(out, *NewRetweetViewDTO(&vs[i]))
	}
	return out
}
