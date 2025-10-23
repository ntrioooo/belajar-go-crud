package services

import (
	"context"
	"errors"
	"strings"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"
)

type retweetService struct {
	retweets ports.RetweetRepository
	posts    ports.PostRepository
}

func NewRetweetService(retweets ports.RetweetRepository, posts ports.PostRepository) ports.RetweetService {
	return &retweetService{retweets: retweets, posts: posts}
}

// Toggle: jika sudah ada retweet user utk post tsb → hapus (unretweet).
// Jika belum ada → buat. Jika quoteBody != nil && != "" → simpan sebagai quote.
func (s *retweetService) Toggle(ctx context.Context, userID, originalPostID uint, quoteBody *string) (bool, int64, error) {
	if userID == 0 {
		return false, 0, errors.New("unauthorized")
	}
	// pastikan post exist
	if p, _ := s.posts.FindByID(ctx, originalPostID); p == nil {
		return false, 0, errors.New("post not found")
	}

	exist, _ := s.retweets.FindByUserAndPost(ctx, userID, originalPostID)
	if exist != nil {
		// unretweet
		if err := s.retweets.Delete(ctx, userID, originalPostID); err != nil {
			return false, 0, err
		}
		cnt, _ := s.retweets.CountByPost(ctx, originalPostID)
		return false, cnt, nil
	}

	var q *string
	if quoteBody != nil {
		val := strings.TrimSpace(*quoteBody)
		if val != "" {
			q = &val
		}
	}
	rt := &domain.Retweet{
		UserID:         userID,
		OriginalPostID: originalPostID,
		QuoteBody:      q,
	}
	if err := s.retweets.Create(ctx, rt); err != nil {
		return false, 0, err
	}
	cnt, _ := s.retweets.CountByPost(ctx, originalPostID)
	return true, cnt, nil
}

func (s *retweetService) ListByUser(ctx context.Context, userID uint) ([]domain.RetweetView, error) {
	items, err := s.retweets.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	out := make([]domain.RetweetView, 0, len(items))

	// Ambil ringkasan original post
	for i := range items {
		rt := items[i]
		p, _ := s.posts.FindByID(ctx, rt.OriginalPostID)
		if p == nil {
			// post sudah dihapus? lewati dengan aman
			continue
		}
		authorMap, _ := s.posts.BatchAuthorUsernames(ctx, []uint{p.UserID})
		catMap, _ := s.posts.BatchCategoryNames(ctx, []uint{p.CategoryID})
		out = append(out, domain.RetweetView{
			Retweet: rt,
			Original: domain.PostSummary{
				ID:             p.ID,
				Title:          p.Title,
				Body:           p.Body,
				AuthorUsername: authorMap[p.UserID],
				CategoryName:   catMap[p.CategoryID],
			},
		})
	}
	return out, nil
}
