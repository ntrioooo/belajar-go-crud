package dto

import "belajar-go/internal/core/domain"

type ProfileDTO struct {
	UserID       uint          `json:"user_id"`
	Username     string        `json:"username"`
	PostCount    int64         `json:"post_count"`
	RetweetCount int64         `json:"retweet_count"`
	Posts        []PostViewDTO `json:"posts"`
}

func NewProfileDTO(v *domain.ProfileView) *ProfileDTO {
	if v == nil {
		return nil
	}
	return &ProfileDTO{
		UserID:       v.UserID,
		Username:     v.Username,
		PostCount:    v.PostCount,
		RetweetCount: v.RetweetCount,
		Posts:        NewPostViewDTOs(v.Posts),
	}
}
